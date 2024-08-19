#!/bin/bash
set -e

echo "Running tests for GPT-like Chat Application..."

# Ensure we're in the project root
cd "$(dirname "$0")/.."

# Run all tests
go test -v ./...

echo "All tests completed."