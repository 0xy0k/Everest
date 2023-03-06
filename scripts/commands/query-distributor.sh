#!/bin/bash

# query collected fees so far
tsukid query distributor fees-collected

# query fees treasury 
tsukid query distributor fees-treasury

# query snapshot period
tsukid query distributor snapshot-period

# query validator performance
tsukid query distributor snapshot-period-performance tsukivaloper177endc8f4tnl44q0x2qct6g9l5af6nm2vlzxd6

# query year start snapshot to check inflation from snapshot time
tsukid query distributor year-start-snapshot