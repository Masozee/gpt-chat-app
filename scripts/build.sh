#!/bin/bash
set -e

echo "Building GPT-like Chat Application..."

# Ensure we're in the project root
cd "$(dirname "$0")/.."

# Run tests
go test ./...

# Build the application
go build -o gpt-chat-app ./cmd/server

echo "Build complete. Binary is './gpt-chat-app'"