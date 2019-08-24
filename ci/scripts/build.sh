#!/usr/bin/env sh

set -eu

BUILDCTL_ARGS=${BUILDCTL_ARGS:-}
CACHE_BASE=${CACHE_BASE:-/tmp/buildkit-caches}
TARGET=${1:-}

if [ -n "${TARGET}" ]; then
    BUILDCTL_ARGS="${BUILDCTL_ARGS} --opt target=${TARGET}"

    if [ -n "${CACHE_BASE}" ]; then
        mkdir -p ${CACHE_BASE}
        BUILDCTL_ARGS="${BUILDCTL_ARGS} --export-cache type=local,mode=max,dest=${CACHE_BASE}/${TARGET}"
    fi
fi

if [ -n "${CACHE_BASE}" ]; then
    for d in $(find ${CACHE_BASE} -maxdepth 1 -mindepth 1 -type d); do
        if [ -f ${d}/index.json ]; then
            BUILDCTL_ARGS="${BUILDCTL_ARGS} --import-cache type=local,src=${d}"
        fi
    done
fi

buildctl build --frontend dockerfile.v0 --local context=. --local dockerfile=. $BUILDCTL_ARGS
