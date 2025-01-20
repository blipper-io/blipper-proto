#!/bin/bash
set -e

# Get version from git tag
VERSION=$(git describe --tags --abbrev=0)
VERSION_NO_V=${VERSION#v}

# Publish Go package
cd gen/go
GOPROXY=proxy.golang.org go list -m github.com/blipper-io/blipper-proto@${VERSION}

# Publish NuGet package
cd ../csharp/Blipper.Proto
dotnet pack -c Release -p:Version=${VERSION_NO_V} -p:RepositoryUrl=https://github.com/blipper-io/blipper-proto
dotnet nuget push "bin/Release/*.nupkg" --source "github"