#!/bin/bash

set -Eeuo pipefail

function cleanup {      
  echo bye
}
trap cleanup EXIT

[ -z ${1-} ] && { echo "use build.sh VersionCore [PreRelease]"; exit 1; }

VersionCore=$1
PreRelease=${2-}
[ -z ${PreRelease-} ] && SemVer=$VersionCore || SemVer=$VersionCore-$PreRelease

[ -d "./.build" ] || mkdir ./.build

pushd . > /dev/null
cd cli
env GOOS=linux GOARCH=amd64 go build -o ../.build/ce
popd > /dev/null

pushd .build > /dev/null
tar -czvf ce_linux_amd64.tar.gz ce
popd > /dev/null

cp install/install.sh .build

