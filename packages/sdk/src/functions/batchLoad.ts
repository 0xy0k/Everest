import { BigNumber } from '@ethersproject/bignumber';
import { AddressZero } from '@ethersproject/constants';
import { formatUnits } from '@ethersproject/units';
import { Call } from '@hovoh/ethcall';

import { EverestErrorCode } from '../constants';
import { CHIEF_ADDRESS, EVEREST_ORACLE_ADDRESS } from '../constants/addresses';
import { LENDING_PROVIDERS } from '../constants/lending-providers';
import {
  Address,
  EverestError,
  EverestResultError,
  EverestResultSuccess,
} from '../entities';
import { AbstractVault } from '../entities/abstract/AbstractVault';
import { BorrowingVault } from '../entities/BorrowingVault';
import { Chain } from '../entities/Chain';
import { VaultType } from '../enums';
import { EverestResult, EverestResultPromise, VaultWithFinancials } from '../types';
import {
  Chief__factory,
  EverestOracle__factory,
  ILendingProvider__factory,
} from '../types/contracts';
import { ChiefMulticall } from '../types/contracts/src/Chief';
import { EverestOracleMulticall } from '../types/contracts/src/EverestOracle';

// number of details calls per vault
const N_CALLS = 10;

type Detail = BigNumber | string | string[];
type Rate = BigNumber;

// rates are with 27 decimals
const rateToFloat = (n: BigNumber) =>
  parseFloat(formatUnits(n.toString(), 27)) * 100;

const getDetailsCalls = (
  v: BorrowingVault,
  account: Address | undefined,
  oracle: EverestOracleMulticall,
  chief: ChiefMulticall
): EverestResult<Call<Detail>[]> => {
  if (!v.multicallContract) {
    return new EverestResultError('BorrowingVault multicallContract not set!');
  }

  return new EverestResultSuccess([
    v.multicallContract.maxLtv() as Call<BigNumber>,
    v.multicallContract.liqRatio() as Call<BigNumber>,
    v.multicallContract.name() as Call<string>,
    v.multicallContract.activeProvider() as Call<string>,
    v.multicallContract.getProviders() as Call<string[]>,
    chief.vaultSafetyRating(v.address.value),
    v.multicallContract.balanceOfAsset(
      account?.value ?? AddressZero
    ) as Call<BigNumber>,
    v.multicallContract?.balanceOfDebt(
      account?.value ?? AddressZero
    ) as Call<BigNumber>,
    oracle.getPriceOf(
      AddressZero,
      v.collateral.address.value,
      v.collateral.decimals
    ),
    oracle.getPriceOf(AddressZero, v.debt.address.value, v.debt.decimals),
  ]);
};

const getProvidersCalls = (v: BorrowingVault): EverestResult<Call<Rate>[]> => {
  if (!v.allProviders) {
    return new EverestResultError('BorrowingVault allProviders not set!');
  }

  return new EverestResultSuccess(
    v.allProviders
      .map((addr) => [
        ILendingProvider__factory.multicall(addr).getDepositRateFor(
          v.address.value
        ),
        ILendingProvider__factory.multicall(addr).getBorrowRateFor(
          v.address.value
        ),
      ])
      // flatten [][] to []
      .reduce((acc, b) => acc.concat(...b), [])
  );
};

const setResults = (
  v: BorrowingVault,
  detailsBatch: Detail[],
  rates: Rate[]
): EverestResult<VaultWithFinancials> => {
  if (!v.activeProvider || !v.allProviders) {
    return new EverestResultError(
      'BorrowingVault activeProvider and allProviders not set!'
    );
  }
  const apIndex = v.allProviders.findIndex((addr) => v.activeProvider === addr);
  const providers = v.allProviders.map((addr, i) => {
    return {
      ...LENDING_PROVIDERS[v.chainId][addr],
      depositAprBase: rateToFloat(rates[2 * i]),
      borrowAprBase: rateToFloat(rates[2 * i + 1]),
    };
  });
  const nonActives = providers.filter((_, i) => i !== apIndex);
  const active = providers[apIndex];
  return new EverestResultSuccess({
    vault: v,
    depositBalance: detailsBatch[6] as BigNumber,
    borrowBalance: detailsBatch[7] as BigNumber,
    collateralPriceUSD: detailsBatch[8] as BigNumber,
    debtPriceUSD: detailsBatch[9] as BigNumber,
    allProviders: [active, ...nonActives],
    activeProvider: active,
  });
};

export async function batchLoad(
  type: VaultType,
  vaults: AbstractVault[],
  account: Address | undefined,
  chain: Chain
): EverestResultPromise<VaultWithFinancials[]> {
  if (!chain.connection) {
    return new EverestResultError('Chain connection not set!', EverestErrorCode.SDK, {
      chainId: chain.chainId,
    });
  }
  // TODO: Check that type matches vaults?
  if (vaults.find((v) => v.chainId !== chain.chainId)) {
    return new EverestResultError(
      'Vault from a different chain!',
      EverestErrorCode.SDK,
      {
        chainId: chain.chainId,
      }
    );
  }

  return type === VaultType.BORROW
    ? _borrowingBatchLoad(vaults as BorrowingVault[], account, chain)
    : _borrowingBatchLoad(vaults as BorrowingVault[], account, chain); // TODO:
}

async function _borrowingBatchLoad(
  vaults: BorrowingVault[],
  account: Address | undefined,
  chain: Chain
): EverestResultPromise<VaultWithFinancials[]> {
  if (!chain.connection) {
    return new EverestResultError('Connection not set!');
  }
  try {
    const { multicallRpcProvider } = chain.connection;
    const oracle = EverestOracle__factory.multicall(
      EVEREST_ORACLE_ADDRESS[chain.chainId].value
    );
    const chief = Chief__factory.multicall(CHIEF_ADDRESS[chain.chainId].value);

    const batchResult = vaults.map((v) =>
      getDetailsCalls(v, account, oracle, chief)
    );
    let error = batchResult.find((r): r is EverestResultError => !r.success);
    if (error)
      return new EverestResultError(error.error.message, error.error.code);

    const detailsBatch = (
      batchResult as EverestResultSuccess<Call<Detail>[]>[]
    ).map((r) => r.data);

    const detailsBatchResults = await multicallRpcProvider.all(
      // flatten [][] to []
      detailsBatch.reduce((acc, b) => acc.concat(...b), [])
    );

    vaults.forEach((v, i) => {
      const maxLtv = detailsBatchResults[N_CALLS * i] as BigNumber;
      const liqRatio = detailsBatchResults[N_CALLS * i + 1] as BigNumber;
      const name = detailsBatchResults[N_CALLS * i + 2] as string;
      const activeProvider = detailsBatchResults[N_CALLS * i + 3] as string;
      const allProviders = detailsBatchResults[N_CALLS * i + 4] as string[];
      const safetyRating = detailsBatchResults[N_CALLS * i + 5] as BigNumber;
      v.setBorrowingPreLoads(
        maxLtv,
        liqRatio,
        safetyRating,
        name,
        activeProvider,
        allProviders
      );
    });

    const ratesResult = vaults.map((v) => getProvidersCalls(v));
    error = ratesResult.find((r): r is EverestResultError => !r.success);
    if (error)
      return new EverestResultError(error.error.message, error.error.code);

    const ratesBatch = (ratesResult as EverestResultSuccess<Call<Rate>[]>[]).map(
      (r) => r.data
    );

    // Every vault has a different amount of lending providers.
    // We can't use the same mechanics as for detailsBatch
    // where every vault has a fixed number of attributes.
    // We have to pass a flattened array of calls and
    // that's why for the rates we need to set offsets and length for each vault.
    let memo = 0;
    const offsets: { offset: number; len: number }[] = ratesBatch.map(
      (batch: Call<Rate>[]) => {
        const o = { offset: memo, len: batch.length };
        memo += batch.length;
        return o;
      }
    );
    const ratesBatchResults = await multicallRpcProvider.all(
      // flatten [][] to []
      ratesBatch.reduce((acc, v) => acc.concat(...v), [])
    );

    const result = vaults.map((v, i) => {
      const details = detailsBatchResults.slice(
        N_CALLS * i,
        N_CALLS * i + N_CALLS
      );
      const o = offsets[i];
      const rates = ratesBatchResults.slice(o.offset, o.offset + o.len);
      return setResults(v, details, rates);
    });
    error = result.find((r): r is EverestResultError => !r.success);
    if (error)
      return new EverestResultError(error.error.message, error.error.code);
    const data = (result as EverestResultSuccess<VaultWithFinancials>[]).map(
      (r) => r.data
    );

    return new EverestResultSuccess(data);
  } catch (e: unknown) {
    const message = EverestError.messageFromUnknownError(e);
    return new EverestResultError(message);
  }
}
