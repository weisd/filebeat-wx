#!/usr/bin/env bash

moduleName=$(go list -m)
commit=$(git rev-parse --short HEAD)
branch=$(git rev-parse --abbrev-ref HEAD)
buildTime=$(date +%Y%m%d%H%M)

versionPkgName="version"

flags="-X '${moduleName}/${versionPkgName}.BuildGitBranch=${branch}' -X '${moduleName}/${versionPkgName}.BuildGitCommit=${commit}'  -X '${moduleName}/${versionPkgName}.BuildDateTime=${buildTime}' -X '${moduleName}/${versionPkgName}.BuildPackageModule=${moduleName}' "

program=$(basename ${moduleName})

#echo "$program"
#
#echo "${flags}"
#
#echo "\$(pwd)"
#

GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)

if [ "$1" == "linux" ]; then
    GOARCH="amd64"
    GOOS="linux"
fi

#echo ">> GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$program -ldflags "$flags" -v -mod=vendor  main.go"

GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$program -ldflags "$flags" -v -mod=vendor  main.go