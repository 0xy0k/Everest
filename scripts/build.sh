#!/usr/bin/env bash
set -e
set -x
. /etc/profile

LOCAL_PLATFORM="$(uname)" && LOCAL_PLATFORM="$(echo "$LOCAL_PLATFORM" |  tr '[:upper:]' '[:lower:]' )"
LOCAL_ARCH=$(([[ "$(uname -m)" == *"arm"* ]] || [[ "$(uname -m)" == *"aarch"* ]]) && echo "arm64" || echo "amd64")
LOCAL_OUT="${GOBIN}/tsukid"

PLATFORM="$1" && [ -z "$PLATFORM" ] && PLATFORM="$LOCAL_PLATFORM"
ARCH="$2" && [ -z "$ARCH" ] && ARCH="$LOCAL_ARCH"
OUTPUT="$3" && [ -z "$OUTPUT" ] && OUTPUT="$LOCAL_OUT"

CONSTANS_FILE=./types/constants.go
VERSION=$(grep -Fn -m 1 'TsukiVersion ' $CONSTANS_FILE | rev | cut -d "=" -f1 | rev | xargs | tr -dc '[:alnum:]\-\.' || echo '')
($(isNullOrEmpty "$VERSION")) && ( echoErr "ERROR: TsukiVersion was NOT found in contants '$CONSTANS_FILE' !" && sleep 5 && exit 1 )

COMMIT=$(git log -1 --format='%H') && \
ldfName="-X github.com/cosmos/cosmos-sdk/version.Name=tsuki" && \
ldfAppName="-X github.com/cosmos/cosmos-sdk/version.AppName=tsukid" && \
ldfVersion="-X github.com/cosmos/cosmos-sdk/version.Version=$VERSION" && \
ldfCommit="-X github.com/cosmos/cosmos-sdk/version.Commit=$COMMIT" && \
ldfBuildTags="-X github.com/cosmos/cosmos-sdk/version.BuildTags=${PLATFORM},${ARCH}"

rm -fv "$OUTPUT" || echo "ERROR: Failed to wipe old tsukid binary"

go mod tidy
GO111MODULE=on go mod verify
env GOOS=$PLATFORM GOARCH=$ARCH go build -ldflags "${ldfName} ${ldfAppName} ${ldfVersion} ${ldfCommit} ${ldfBuildTags}" -o "$OUTPUT" ./cmd/tsukid

( [ "$PLATFORM" == "$LOCAL_PLATFORM" ] && [ "$ARCH" == "$LOCAL_ARCH" ] && [ -f $OUTPUT ] ) && \
    echoInfo "INFO: Sucessfully built TSUKI $($OUTPUT version)" || echoInfo "INFO: Sucessfully built TSUKI to '$OUTPUT'"
