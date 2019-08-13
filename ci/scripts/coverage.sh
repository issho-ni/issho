#!/usr/bin/env bash

set -eu

CODECOV_TOKEN=${CODECOV_TOKEN:-}

if [ ! -n "${CODECOV_TOKEN}" ]; then
    echo "required parameter \$CODECOV_TOKEN missing!"
    exit 1
fi

bash <(curl -s https://codecov.io/bash) -Z -f coverage.out
