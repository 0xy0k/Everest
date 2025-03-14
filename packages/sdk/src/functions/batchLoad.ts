import { BigNumber } from '@ethersproject/bignumber';
import { AddressZero } from '@ethersproject/constants';
import { formatUnits } from '@ethersproject/units';
import { Call, IMulticallProvider } from '@hovoh/ethcall';

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
import { LendingVault } from '../entities/LendingVault';
import { ChainId } from '../enums';
import { EverestResult, EverestResultPromise, VaultWithFinancials } from '../types';
import {
  Chief__factory,
  EverestOracle__factory,
  ILendingProvider__factory,
} from '../types/contracts';
import { ChiefMulticall } from '../types/contracts/src/Chief';
import { EverestOracleMulticall } from '../types/contracts/src/EverestOracle';

// number of details calls per vault
const N_CALLS = 11;

type Detail = BigNumber | string | string[];
type Rate = BigNumber;

// rates are with 27 decimals
const rateToFloat = (n: BigNumber) =>
  parseFloat(formatUnits(n.toString(), 27)) * 100;

const getDetailsCalls = (
  v: AbstractVault,
  account: Address | undefined,
  oracle: EverestOracleMulticall,
  chief: ChiefMulticall
): EverestResult<Call<Detail>[]> => {
  if (!v.multicallContract) {
    return new EverestResultError('BorrowingVault multicallContract not set!');
  }

  const empty: Call = {
    contract: {
      address: AddressZero,
    },
    name: '',
    inputs: [],
    outputs: [],
    params: [],
  };

  if (v instanceof BorrowingVault) {
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
      v.multicallContract.VERSION() as Call<string>,
    ]);
  } else {
    // pass an empty call for methods that don't exist on LendingVault
    // to preserve the indexing
    return new EverestResultSuccess([
      empty,
      empty,
      v.multicallContract.name() as Call<string>,
      v.multicallContract.activeProvider() as Call<string>,
      v.multicallContract.getProviders() as Call<string[]>,
      chief.vaultSafetyRating(v.address.value),
      v.multicallContract.balanceOfAsset(
        account?.value ?? AddressZero
      ) as Call<BigNumber>,
      empty,
      oracle.getPriceOf(
        AddressZero,
        v.collateral.address.value,
        v.collateral.decimals
      ),
      empty,
      v.multicallContract.VERSION() as Call<string>,
    ]);
  }
};

const getProvidersCalls = (v: AbstractVault): EverestResult<Call<Rate>[]> => {
  if (!v.allProviders) {
    return new EverestResultError('Vault allProviders not set!');
  }

  let rates;

  if (v instanceof BorrowingVault) {
    rates = v.allProviders.map((addr) => [
      ILendingProvider__factory.multicall(addr).getDepositRateFor(
        v.address.value
      ),
      ILendingProvider__factory.multicall(addr).getBorrowRateFor(
        v.address.value
      ),
    ]);
  } else {
    rates = v.allProviders.map((addr) => [
      ILendingProvider__factory.multicall(addr).getDepositRateFor(
        v.address.value
      ),
    ]);
  }

  // flatten [][] to []
  return new EverestResultSuccess(rates.reduce((acc, b) => acc.concat(...b), []));
};

const setResults = (
  v: AbstractVault,
  detailsBatch: (Detail | null)[],
  rates: Rate[]
): EverestResult<VaultWithFinancials> => {
  if (!v.activeProvider || !v.allProviders) {
    return new EverestResultError(
      'BorrowingVault activeProvider and allProviders not set!'
    );
  }
  const apIndex = v.allProviders.findIndex((addr) => v.activeProvider === addr);
  const providers = v.allProviders.map((addr, i) => {
    if (v instanceof BorrowingVault) {
      return {
        ...LENDING_PROVIDERS[v.chainId][addr],
        depositAprBase: rateToFloat(rates[2 * i]),
        borrowAprBase: rateToFloat(rates[2 * i + 1]),
      };
    } else {
      return {
        ...LENDING_PROVIDERS[v.chainId][addr],
        depositAprBase: rateToFloat(rates[i]),
      };
    }
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

const setVaultsPreLoads = (
  vaults: AbstractVault[],
  detailsBatchResults: (Detail | null)[]
) => {
  vaults.forEach((v, i) => {
    const maxLtv = detailsBatchResults[N_CALLS * i] as BigNumber;
    const liqRatio = detailsBatchResults[N_CALLS * i + 1] as BigNumber;
    const name = detailsBatchResults[N_CALLS * i + 2] as string;
    const activeProvider = detailsBatchResults[N_CALLS * i + 3] as string;
    const allProviders = detailsBatchResults[N_CALLS * i + 4] as string[];
    const safetyRating = detailsBatchResults[N_CALLS * i + 5] as BigNumber;
    const version = detailsBatchResults[N_CALLS * i + 10] as string;
    if (v instanceof BorrowingVault) {
      v.setPreLoads(
        maxLtv,
        liqRatio,
        safetyRating,
        name,
        activeProvider,
        allProviders,
        version
      );
    } else if (v instanceof LendingVault) {
      v.setPreLoads(safetyRating, name, activeProvider, allProviders, version);
    }
  });
};

const doBatchLoad = async (
  vaults: AbstractVault[],
  account: Address | undefined,
  multicall: IMulticallProvider,
  chain: Chain
): EverestResultPromise<VaultWithFinancials[]> => {
  const oracle = EverestOracle__factory.multicall(
    EVEREST_ORACLE_ADDRESS[chain.chainId].value
  );
  const chief = Chief__factory.multicall(CHIEF_ADDRESS[chain.chainId].value);

  const batchResult = vaults.map((v) =>
    getDetailsCalls(v, account, oracle, chief)
  );
  let error = batchResult.find((r): r is EverestResultError => !r.success);
  if (error) return new EverestResultError(error.error.message, error.error.code);

  const detailsBatch = (batchResult as EverestResultSuccess<Call<Detail>[]>[]).map(
    (r) => r.data
  );
  const detailsBatchResults = await multicall.tryAll(
    // flatten [][] to []
    detailsBatch.reduce((acc, b) => acc.concat(...b), [])
  );

  setVaultsPreLoads(vaults, detailsBatchResults);

  const ratesResult = vaults.map((v) => getProvidersCalls(v));
  error = ratesResult.find((r): r is EverestResultError => !r.success);
  if (error) return new EverestResultError(error.error.message, error.error.code);

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
  const ratesBatchResults = await multicall.all(
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
  if (error) return new EverestResultError(error.error.message, error.error.code);
  const data = (result as EverestResultSuccess<VaultWithFinancials>[]).map(
    (r) => r.data
  );

  return new EverestResultSuccess(data);
};

export async function multiBatchLoad(
  vaults: AbstractVault[],
  account: Address | undefined
): EverestResultPromise<VaultWithFinancials[]> {
  try {
    // Group vaults by chainId
    const vaultsByChainId: Record<
      ChainId,
      { vaults: AbstractVault[]; chain: Chain }
    > = Object.create(null);

    for (const vault of vaults) {
      const chainId = vault.chainId;
      if (!vaultsByChainId[chainId]) {
        vaultsByChainId[chainId] = { vaults: [], chain: vault.chain };
      }
      vaultsByChainId[chainId].vaults.push(vault);
    }

    // Make batchLoad calls for each group
    const results = await Promise.all(
      Object.values(vaultsByChainId).map(({ vaults, chain }) =>
        batchLoad(vaults, account, chain)
      )
    );

    // Combine all successful results into a single array
    const combinedResults: VaultWithFinancials[] = results.reduce(
      (acc: VaultWithFinancials[], result) => {
        if (result.success) {
          return [...acc, ...result.data];
        } else {
          // Ignore errors for the time being since we need to handle them
          return acc;
        }
      },
      []
    );
    if (combinedResults.length === 0) {
      throw new Error('No results found!');
    }
    return new EverestResultSuccess(combinedResults);
  } catch (error) {
    const message = EverestError.messageFromUnknownError(error);
    return new EverestResultError(message);
  }
}

export async function batchLoad(
  vaults: AbstractVault[],
  account: Address | undefined,
  chain: Chain
): EverestResultPromise<VaultWithFinancials[]> {
  if (!chain.connection) {
    return new EverestResultError('Chain connection not set!', EverestErrorCode.SDK, {
      chainId: chain.chainId,
    });
  }
  if (
    vaults.find(
      (v) => v instanceof BorrowingVault && v.chainId !== chain.chainId
    )
  ) {
    return new EverestResultError(
      'Borrowing vault from a different chain!',
      EverestErrorCode.SDK,
      {
        chainId: chain.chainId,
      }
    );
  }

  const { multicallRpcProvider: multicall } = chain.connection;

  try {
    return doBatchLoad(vaults, account, multicall, chain);
  } catch (e: unknown) {
    const message = EverestError.messageFromUnknownError(e);
    return new EverestResultError(message);
  }
}
