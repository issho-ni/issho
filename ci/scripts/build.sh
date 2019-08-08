#!/usr/bin/env sh

set -eu

if [ -n "$DOCKER_HUB_USERNAME" ] && [ -n "$DOCKER_HUB_PASSWORD" ]; then
    auth=$(printf "%s:%s" $DOCKER_HUB_USERNAME $DOCKER_HUB_PASSWORD | base64 | tr -d "\n")
    mkdir ~/.docker
    echo "{\"auths\":{\"https://index.docker.io/v1/\":{\"auth\":\"${auth}\"}}}" > ~/.docker/config.json
fi

exec build
