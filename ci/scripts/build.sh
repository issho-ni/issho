#!/usr/bin/env sh

set -eu

auth=$(printf "%s:%s" "_json_key" "${GCR_JSON_KEY}" | base64 | tr -d "\n")
mkdir ~/.docker
cat <<EOF > ~/.docker/config.json
{"auths":{"https://gcr.io/v2/":{"auth":"${auth}"}}}
EOF

build
