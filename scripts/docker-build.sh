#!/bin/bash
set -euo pipefail

# Build Docker image for specified architecture
# Usage: ./scripts/docker-build.sh <arch> <version>

if [[ $# -ne 2 ]]; then
  echo "Usage: $0 <arch> <version>" >&2
  echo "  arch: arm64 or amd64" >&2
  echo "  version: version tag" >&2
  exit 1
fi

ARCH="$1"
VERSION="$2"

case "${ARCH}" in
  arm64) ;;
  amd64) ;;
  *)
    echo "Error: Unsupported architecture '${ARCH}'. Use 'arm64' or 'amd64'." >&2
    exit 1
    ;;
esac

case "$(uname)" in
  Darwin) DATE='gdate' ;;
  *) DATE='date' ;;
esac

rm -f data-*.tar.bz2
BZIP2=--best tar -cjf "data-$(TZ=UTC "${DATE}" +%Y%m%dT%H%M%SZ).tar.bz2" \
  data/person.d \
  data/workbench0/ \
  data/workbench1/ \
  data/workbench2/ \
  data/workbench3/ \
  data/workbench4/ \
  data/workbench5/ \
  data/workbench6/ \
  data/workbench7/ \
  data/workbench8/ \
  data/workbench9/

# Build with architecture-specific tags
docker buildx build \
  --build-arg ARCH="${ARCH}" \
  --build-arg VERSION="${VERSION}" \
  --platform="linux/${ARCH}" \
  --tag="federation-1999:latest" \
  --tag="federation-1999:${VERSION/+/-}" \
  --file=deployments/docker/Dockerfile \
  .
