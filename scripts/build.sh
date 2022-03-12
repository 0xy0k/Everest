#!/usr/bin/env bash
set -e
set -x
. /etc/profile


CONSTANS_FILE=./types/constants.go
TSUKI_VER=$(grep -Fn -m 1 'TsukiVersion ' $CONSTANS_FILE | rev | cut -d "=" -f1 | rev | xargs | tr -dc '[:alnum:]\-\.' || echo '')
($(isNullOrEmpty "$TSUKI_VER")) && ( echoErr "ERROR: TsukiVersion was NOT found in contants '$CONSTANS_FILE' !" && sleep 5 && exit 1 )

COMMIT=$(git log -1 --format='%H') && \
ldfName="-X github.com/cosmos/cosmos-sdk/version.Name=tsuki" && \
ldfAppName="-X github.com/cosmos/cosmos-sdk/version.AppName=tsukid" && \
ldfVersion="-X github.com/cosmos/cosmos-sdk/version.Version=$TSUKI_VER" && \
ldfCommit="-X github.com/cosmos/cosmos-sdk/version.Name=$COMMIT"

rm -rfv "${GOBIN}/tsukid" || echo "ERROR: Failed to wipe old tsukid binary"

go mod tidy
GO111MODULE=on go mod verify
go build -ldflags "${ldfName} ${ldfAppName} ${ldfVersion} ${ldfCommit}" -o "${GOBIN}/tsukid" ./cmd/tsukid
echoInfo "INFO: Sucessfully intalled TSUKI $(tsukid version)"
