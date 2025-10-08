#!/bin/bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 VERSION" >&2
  exit 1
fi
readonly VERSION="$1"

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

docker buildx build \
  --file deployments/docker/Dockerfile \
  --platform linux/amd64 \
  --tag "federation-1999:latest" \
  --tag "federation-1999:${VERSION/+/-}" \
  .
