#!/bin/sh

set -eu

echo "Starting tests..."
exec go list ./... | grep -v "mock" | xargs go test
