#!/bin/sh

set -eu

echo "Starting tests..."

go list ./... | grep -v "mock" | xargs go test -coverprofile=coverage.out -covermode=atomic
