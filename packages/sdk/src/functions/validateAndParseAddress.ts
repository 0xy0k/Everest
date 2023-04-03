import { getAddress } from '@ethersproject/address';

import { EverestResultError, EverestResultSuccess } from '../entities';
import { EverestResult } from '../types';

// warns if addresses are not checksummed
export function validateAndParseAddress(address: string): EverestResult<string> {
  try {
    const checksummedAddress = getAddress(address);
    if (address !== checksummedAddress)
      console.warn(`${address} is not checksummed.`);
    //warning(address === checksummedAddress, `${address} is not checksummed.`);
    return new EverestResultSuccess(checksummedAddress);
  } catch (error) {
    return new EverestResultError(`${address} is not a valid address.`);
  }
}
