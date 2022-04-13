#!/usr/bin/env bash
set -e
set -x
. /etc/profile
. ./scripts/tsuki-env.sh

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

UTILS_VER=$(bashUtilsVersion 2> /dev/null || echo "")
UTILS_OLD_VER="false" && [[ $(versionToNumber "$UTILS_VER" || echo "0") -ge $(versionToNumber "v0.1.2.3" || echo "1") ]] || UTILS_OLD_VER="true" 

# Installing utils is essential to simplify the setup steps
if [ "$UTILS_OLD_VER" == "true" ] ; then
    echo "INFO: TSUKI utils were NOT installed on the system, setting up..." && sleep 2
    TOOLS_VERSION="v0.0.12.4" && mkdir -p /usr/keys && FILE_NAME="bash-utils.sh" && \
     if [ -z "$TSUKI_COSIGN_PUB" ] ; then TSUKI_COSIGN_PUB=/usr/keys/tsuki-cosign.pub ; fi && \
     echo -e "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE/IrzBQYeMwvKa44/DF/HB7XDpnE+\nf+mU9F/Qbfq25bBWV2+NlYMJv3KvKHNtu3Jknt6yizZjUV4b8WGfKBzFYw==\n-----END PUBLIC KEY-----" > $TSUKI_COSIGN_PUB && \
     wget "https://github.com/TsukiCore/tools/releases/download/$TOOLS_VERSION/${FILE_NAME}" -O ./$FILE_NAME && \
     wget "https://github.com/TsukiCore/tools/releases/download/$TOOLS_VERSION/${FILE_NAME}.sig" -O ./${FILE_NAME}.sig && \
     cosign verify-blob --key="$TSUKI_COSIGN_PUB" --signature=./${FILE_NAME}.sig ./$FILE_NAME && \
     chmod -v 555 ./$FILE_NAME && ./$FILE_NAME bashUtilsSetup "/var/tsukiglob" && . /etc/profile && \
     echoInfo "Installed bash-utils $(bashUtilsVersion)"
else
    echoInfo "INFO: TSUKI utils are up to date, latest version $UTILS_VER" && sleep 2
fi

echoInfo "INFO: Environment cleanup...."
systemctl2 stop tsuki || echoWarn "WARNING: tsuki service was NOT running or could NOT be stopped"

kill -9 $(sudo lsof -t -i:9090) || echoWarn "WARNING: Nothing running on port 9090, or failed to kill processes"
kill -9 $(sudo lsof -t -i:6060) || echoWarn "WARNING: Nothing running on port 6060, or failed to kill processes"
kill -9 $(sudo lsof -t -i:26656) || echoWarn "WARNING: Nothing running on port 26656, or failed to kill processes"
kill -9 $(sudo lsof -t -i:26657) || echoWarn "WARNING: Nothing running on port 26657, or failed to kill processes"
kill -9 $(sudo lsof -t -i:26658) || echoWarn "WARNING: Nothing running on port 26658, or failed to kill processes"

NETWORK_NAME="localnet-0"
setGlobEnv TSUKID_HOME ~/.tsukid-local
setGlobEnv NETWORK_NAME $NETWORK_NAME
loadGlobEnvs

rm -rfv $TSUKID_HOME 
mkdir -p $TSUKID_HOME

echoInfo "INFO: Starting new network..."

tsukid init --overwrite --chain-id=$NETWORK_NAME "TSUKI TEST LOCAL VALIDATOR NODE" --home=$TSUKID_HOME
tsukid keys add validator --keyring-backend=test --home=$TSUKID_HOME
tsukid keys add faucet --keyring-backend=test --home=$TSUKID_HOME
tsukid add-genesis-account $(showAddress validator) 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=test --home=$TSUKID_HOME
tsukid add-genesis-account $(showAddress faucet) 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=test --home=$TSUKID_HOME
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

if [ "$NETWORK_NAME" != "$NETWORK_STATUS_CHAIN_ID" ] ; then
    echoErr "ERROR: Incorrect chain ID from the status query, expected '$NETWORK_NAME', but got $NETWORK_STATUS_CHAIN_ID"
fi

BLOCK_HEIGHT=$(showBlockHeight)
echoInfo "INFO: Waiting for next block to be produced..." && sleep 10
NEXT_BLOCK_HEIGHT=$(showBlockHeight)

if [ $BLOCK_HEIGHT -ge $NEXT_BLOCK_HEIGHT ] ; then
    echoErr "ERROR: Failed to produce next block height, stuck at $BLOCK_HEIGHT"
fi

echoInfo "INFO: NETWORK-SETUP - Integration Test - END, elapsed: $(prettyTime $(timerSpan $TEST_NAME))"