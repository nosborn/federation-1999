#!/bin/bash
set -euo pipefail

docker kill federation-1999 2>/dev/null || true

exec docker run \
  --mount "type=tmpfs,target=/home/fed/backup" \
  --mount "type=bind,source=${PWD}/data,target=/home/fed/data" \
  --mount "type=tmpfs,target=/home/fed/lock" \
  --mount "type=tmpfs,target=/home/fed/log" \
  --mount "type=tmpfs,target=/home/fed/run" \
  --mount "type=bind,source=${PWD}/test/robobod,target=/home/fed/test/robobod,ro" \
  --name federation-1999 \
  --platform linux/amd64 \
  --publish 23:23 \
  --publish 80:8080 \
  --rm \
  federation-1999:latest
