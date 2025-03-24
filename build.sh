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

go build -o bin/$program -ldflags "$flags" -v -mod=vendor  main.go