#!/usr/bin/env sh

echo "{\"imageLayoutVersion\":\"1.0.0\"}" > cache/oci-layout
tar cf cache/image.tar cache/*
