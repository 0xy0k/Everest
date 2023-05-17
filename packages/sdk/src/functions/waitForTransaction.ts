import { JsonRpcProvider, TransactionReceipt } from '@ethersproject/providers';

import { EverestErrorCode } from '../constants';
import { EverestResultError, EverestResultSuccess } from '../entities';
import { EverestResultPromise } from '../types';

export async function waitForTransaction(
  provider: JsonRpcProvider,
  txHash: string
): EverestResultPromise<TransactionReceipt> {
  const result = await _retry(_waitForTransaction, provider, txHash);
  console.log(result);
  return result;
}

async function _waitForTransaction(
  provider: JsonRpcProvider,
  txHash: string
): EverestResultPromise<TransactionReceipt> {
  try {
    const receipt = await provider.waitForTransaction(txHash);
    return new EverestResultSuccess(receipt);
  } catch (error) {
    return new EverestResultError(
      'Transaction failed on chain', // TODO: message
      EverestErrorCode.TX,
      {
        txHash,
        message: String(error), // TODO:
      }
    );
  }
}

async function _retry<T>(
  fn: (...args: any[]) => EverestResultPromise<T>,
  ...args: any[]
): Promise<EverestResultPromise<T>> {
  try {
    return await fn(...args);
  } catch (error) {
    // TODO: Should we add a delay here?
    // TODO: If tx was for instance reverted, don't even attempt to retry
    return await fn(...args);
  }
}
