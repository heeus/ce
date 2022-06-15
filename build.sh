#!/bin/bash

set -Eeuo pipefail

[ -z ${1-} ] && { echo "use build.sh VersionCore [PreRelease]"; exit 1; }

BuildFolder=".build"

# Process arguments

VersionCore=$1
PreRelease=${2-}
[ -z ${PreRelease-} ] && SemVer=$VersionCore || SemVer=$VersionCore-$PreRelease

# Functions

function cleanup {      
  echo bye
}

function build_os_arch {
  pushd . > /dev/null
  cd cli
  env GOOS=$1 GOARCH=$2 go build -o ../$BuildFolder/ce
  popd > /dev/null

  pushd $BuildFolder > /dev/null
  tar -czvf ce_v${SemVer}_$1_$2.tar.gz ce > /dev/null
  popd > /dev/null
  rm $BuildFolder/ce
}

# End of functions

trap cleanup EXIT

# Cleanup

rm -rf $BuildFolder
[ -d "$BuildFolder" ] || mkdir $BuildFolder

# Build os+arch

build_os_arch linux amd64
build_os_arch linux 386
build_os_arch freebsd amd64
build_os_arch freebsd 386

cp install/install.sh $BuildFolder
