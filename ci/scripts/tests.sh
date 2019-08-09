#!/bin/sh

set -eu

echo "Starting tests..."
exec go test ./...
