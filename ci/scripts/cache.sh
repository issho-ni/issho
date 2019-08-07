#!/usr/bin/env sh

set -eu

echo "{\"imageLayoutVersion\":\"1.0.0\"}" > cache/oci-layout
tar cf cache/image-oci.tar cache/*
/skopeo copy tarball:cache/image-oci.tar docker-archive:cache/image.tar
rm cache/image-oci.tar
