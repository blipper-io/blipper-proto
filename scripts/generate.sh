#!/bin/bash

set -e

# Navigate to repository root
cd "$(dirname "$0")/.."

# Ensure buf CLI is installed
if ! command -v buf &> /dev/null; then
    echo "Buf CLI not found. Please install it from https://buf.build"
    exit 1
fi

echo "Generating Golang & Dotnet code..."
protoc \
--csharp_out=csharp/Blipper.Proto/Generated \
--grpc_out=csharp/Blipper.Proto/Generated \
--plugin=protoc-gen-grpc=$(which grpc_csharp_plugin) \
--proto_path=proto/v1 \
proto/v1/**/*.proto


protoc \
--proto_path=proto/v1 \
--go_out=gen/go/ \
--go-grpc_out=gen/go/ \
proto/v1/**/*.proto

echo "Code generation completed successfully."