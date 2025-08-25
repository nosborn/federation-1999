#!/bin/bash

set -euo pipefail

ARCH=${1:-arm64}
readonly ARCH

case "${ARCH}" in
  amd64 | arm64) ;;
  *)
    echo "Usage: ${0} [amd64|arm64]" >&2
    exit 1
    ;;
esac

docker kill federation-1999 2>/dev/null || true

exec docker run \
  --cap-add NET_ADMIN \
  --mount "type=tmpfs,target=/home/fed/backup" \
  --mount "type=bind,source=${PWD}/data,target=/home/fed/data" \
  --mount "type=tmpfs,target=/home/fed/lock" \
  --mount "type=tmpfs,target=/home/fed/log" \
  --mount "type=tmpfs,target=/home/fed/run" \
  --mount "type=bind,source=${PWD}/test/robobod,target=/home/fed/test/robobod,ro" \
  --mount "type=tmpfs,dst=/run" \
  --mount "type=tmpfs,dst=/tmp" \
  --mount "type=tmpfs,dst=/var/opt/fed" \
  --name federation-1999 \
  --platform "linux/${ARCH}" \
  --publish 23:23 \
  --rm \
  federation-1999:latest
