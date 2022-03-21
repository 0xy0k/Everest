#!/usr/bin/env bash
set -e
set -x
. /etc/profile

echo "INFO: Started local tests in '$PWD'..."
timerStart

echoInfo "INFO: Installing latest tsuki-utils release..."
./scripts/tsuki-utils.sh tsukiUtilsSetup
loadGlobEnvs

echoInfo "INFO: Ensuring correct tsukid version is installed..."
TSUKID_VERSION=$(tsukid version)
TSUKID_EXPECTED_VERSION=$(grep -Fn -m 1 'TsukiVersion ' ./types/constants.go | rev | cut -d "=" -f1 | rev | xargs | tr -dc '[:alnum:]\-\.' || echo '')
RELEASE_EXPECTED_VERSION=$(grep -Fn -m 1 'Release: ' ./RELEASE.md | rev | cut -d ":" -f1 | rev | xargs | tr -dc '[:alnum:]\-\.' || echo '')

[ "$TSUKID_VERSION" != "$TSUKID_EXPECTED_VERSION" ] && \
    echoErr "ERROR: Expected installed tsukid version to be $TSUKID_EXPECTED_VERSION, but got $TSUKID_VERSION, try to make build & install first" && exit 1
[ "$TSUKID_EXPECTED_VERSION" != "$RELEASE_EXPECTED_VERSION" ] && \
    echoErr "ERROR: Expected release version to be $TSUKID_EXPECTED_VERSION, but got $RELEASE_EXPECTED_VERSION, try to make build & install first" && exit 1

echoInfo "INFO: Launching local network..."
./scripts/test-local/network-setup.sh || ( systemctl2 stop tsuki && exit 1 )

echoInfo "INFO: Stopping local network..."
systemctl2 stop tsuki

echoInfo "INFO: Success, all local tests passed, elapsed: $(prettyTime $(timerSpan))"