#!/bin/bash

set -e

# Navigate to repository root
cd "$(dirname "$0")/.."

# Ensure buf CLI is installed
if ! command -v buf &> /dev/null; then
    echo "Buf CLI not found. Please install it from https://buf.build"
    exit 1
fi

echo "Generating Golang code..."
buf generate --template buf.gen.yaml --path proto

echo "Code generation completed successfully."