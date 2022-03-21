#!/usr/bin/env bash
set -e
set -x
. /etc/profile

TEST_NAME="NETWORK-SETUP"
timerStart $TEST_NAME
echoInfo "INFO: $TEST_NAME - Integration Test - START"

echoInfo "INFO: Ensuring essential dependencies are installed & up to date"
SYSCTRL_DESTINATION=/usr/local/bin/systemctl2
if [ ! -f $SYSCTRL_DESTINATION ] ; then
    safeWget /usr/local/bin/systemctl2 \
     https://raw.githubusercontent.com/gdraheim/docker-systemctl-replacement/9cbe1a00eb4bdac6ff05b96ca34ec9ed3d8fc06c/files/docker/systemctl.py \
     "e02e90c6de6cd68062dadcc6a20078c34b19582be0baf93ffa7d41f5ef0a1fdd" && \
    chmod +x $SYSCTRL_DESTINATION && \
    systemctl2 --version
fi

echoInfo "INFO: Environment cleanup...."
systemctl2 stop tsuki || echoWarn "WARNING: tsuki service was NOT running or could NOT be stopped"

setGlobEnv TSUKID_HOME ~/.tsukid-local
loadGlobEnvs

rm -rfv $TSUKID_HOME 
mkdir -p $TSUKID_HOME

echoInfo "INFO: Starting new network..."
CHAIN_ID="localnet-0"
tsukid init --overwrite --chain-id=$CHAIN_ID "TSUKI TEST LOCAL VALIDATOR NODE" --home=$TSUKID_HOME
tsukid keys add validator --keyring-backend=test --home=$TSUKID_HOME
tsukid add-genesis-account $(showAddress validator) 300000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --home=$TSUKID_HOME
tsukid gentx-claim validator --keyring-backend=test --moniker="GENESIS VALIDATOR" --home=$TSUKID_HOME

cat > /etc/systemd/system/tsuki.service << EOL
[Unit]
Description=Local TSUKI Test Network
After=network.target
[Service]
MemorySwapMax=0
Type=simple
User=root
WorkingDirectory=/root
ExecStart=$GOBIN/tsukid start --home=$TSUKID_HOME --trace
Restart=always
RestartSec=5
LimitNOFILE=4096
[Install]
WantedBy=default.target
EOL

systemctl2 enable tsuki 
systemctl2 start tsuki

echoInfo "INFO: Waiting for network to start..." && sleep 3

systemctl2 status tsuki

echoInfo "INFO: Checking network status..."
NETWORK_STATUS_CHAIN_ID=$(showStatus | jq .NodeInfo.network | xargs)

if [ "$CHAIN_ID" != "$NETWORK_STATUS_CHAIN_ID" ] ; then
    echoErr "ERROR: Incorrect chain ID from the status query, expected '$CHAIN_ID', but got $NETWORK_STATUS_CHAIN_ID"
fi

BLOCK_HEIGHT=$(showBlockHeight)
echoInfo "INFO: Waiting for next block to be produced..." && sleep 10
NEXT_BLOCK_HEIGHT=$(showBlockHeight)

if [ $BLOCK_HEIGHT -ge $NEXT_BLOCK_HEIGHT ] ; then
    echoErr "ERROR: Failed to produce next block height, stuck at $BLOCK_HEIGHT"
fi

echoInfo "INFO: NETWORK-SETUP - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"