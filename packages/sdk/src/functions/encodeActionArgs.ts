import { defaultAbiCoder } from '@ethersproject/abi';
import { AddressZero } from '@ethersproject/constants';
import { BigNumber } from 'ethers';

import { EverestResultError, EverestResultSuccess } from '../entities/EverestError';
import { RouterAction } from '../enums';
import { EverestResult, PermitParams, RouterActionParams } from '../types';

const ZERO_BYTES32 =
  '0x0000000000000000000000000000000000000000000000000000000000000000';

function getZeroPermitEncodedArgs(params: PermitParams): string {
  return defaultAbiCoder.encode(
    [
      'address',
      'address',
      'address',
      'uint256',
      'uint256',
      'uint256',
      'bytes32',
      'bytes32',
    ],
    [
      params.vault.value,
      params.owner.value,
      params.receiver.value,
      params.amount.toString(),
      '0',
      '0',
      ZERO_BYTES32,
      ZERO_BYTES32,
    ]
  );
}

export function encodeActionArgs(
  params: RouterActionParams,
  replacePermitWithZero: boolean
): EverestResult<string> {
  let result = '';
  if (
    params.action === RouterAction.DEPOSIT ||
    params.action === RouterAction.PAYBACK
  ) {
    result = defaultAbiCoder.encode(
      ['address', 'uint256', 'address', 'address'],
      [
        params.vault.value,
        params.amount.toString(),
        params.receiver.value,
        params.sender.value,
      ]
    );
  } else if (
    params.action === RouterAction.BORROW ||
    params.action === RouterAction.WITHDRAW
  ) {
    result = defaultAbiCoder.encode(
      ['address', 'uint256', 'address', 'address'],
      [
        params.vault.value,
        params.amount.toString(),
        params.receiver.value,
        params.owner.value,
      ]
    );
  } else if (
    params.action === RouterAction.PERMIT_BORROW ||
    params.action === RouterAction.PERMIT_WITHDRAW
  ) {
    if (replacePermitWithZero) {
      result = getZeroPermitEncodedArgs(params);
    } else if (!(params.deadline && params.v && params.r && params.s)) {
      return new EverestResultError('Missing args in PERMIT_BORROW!');
    } else {
      result = defaultAbiCoder.encode(
        [
          'address',
          'address',
          'address',
          'uint256',
          'uint256',
          'uint8',
          'bytes32',
          'bytes32',
        ],
        [
          params.vault.value,
          params.owner.value,
          params.receiver.value,
          params.amount.toString(),
          params.deadline.toString(),
          params.v.toString(),
          params.r,
          params.s,
        ]
      );
    }
  } else if (params.action === RouterAction.X_TRANSFER) {
    result = defaultAbiCoder.encode(
      ['uint256', 'uint256', 'address', 'uint256', 'address', 'address'],
      [
        params.destDomain,
        params.slippage,
        params.asset.value,
        params.amount.toString(),
        params.receiver.value,
        params.sender.value,
      ]
    );
  } else if (params.action === RouterAction.X_TRANSFER_WITH_CALL) {
    const innerActions = params.innerActions.map(({ action }) =>
      BigNumber.from(action)
    );
    const innerResult = params.innerActions.map((p) =>
      encodeActionArgs(p, replacePermitWithZero)
    );
    const error = innerResult.find((r): r is EverestResultError => !r.success);
    if (error)
      return new EverestResultError(error.error.message, error.error.code);

    const innerArgs: string[] = (
      innerResult as EverestResultSuccess<string>[]
    ).map((r) => r.data);

    const callData = defaultAbiCoder.encode(
      ['uint8[]', 'bytes[]'],
      [innerActions, innerArgs]
    );
    result = defaultAbiCoder.encode(
      ['uint256', 'uint256', 'address', 'uint256', 'address', 'bytes'],
      [
        params.destDomain,
        params.slippage,
        params.asset.isZero ? AddressZero : params.asset.value,
        params.amount.toString(),
        params.sender.value,
        callData,
      ]
    );
  } else if (params.action === RouterAction.DEPOSIT_ETH) {
    result = defaultAbiCoder.encode(['uint256'], [params.amount.toString()]);
  } else if (params.action === RouterAction.WITHDRAW_ETH) {
    result = defaultAbiCoder.encode(
      ['uint256', 'address'],
      [params.amount.toString(), params.receiver.value]
    );
  } else {
    return new EverestResultError('Unsupported action!');
  }

  return new EverestResultSuccess(result);
}
