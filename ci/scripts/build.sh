#!/usr/bin/env sh

set -eu

auth=$(printf "%s:%s" $DOCKER_HUB_USERNAME $DOCKER_HUB_PASSWORD | base64 | tr -d "\n")
mkdir ~/.docker
cat <<EOF > ~/.docker/config.json
{"auths":{"https://index.docker.io/v1/":{"auth":"${auth}"}}}
EOF

build
