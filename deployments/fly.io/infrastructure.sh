#!/usr/bin/env bash
set -euo pipefail

APP_NAME="federation-1999"
BUCKET_NAME="${APP_NAME}"
ORG_NAME="personal"
REGION="iad"
VOLUME_NAME="data"

exists_app() {
  fly apps list --json | jq -e ".[] | select(.Name==\"${APP_NAME}\")" >/dev/null
}

exists_bucket() {
  # fly storage buckets list --json | jq -e ".[] | select(.Name==\"${BUCKET_NAME}\")" >/dev/null
  fly storage list | awk '{print $1}' | grep -qx "${BUCKET_NAME}"
}

exists_ipv4() {
  fly ips list --app "${APP_NAME}" --json | jq -e '.[] | select(.Type=="v4")' >/dev/null
}

exists_volume() {
  fly volumes list --app "${APP_NAME}" --json | jq -e ".[] | select(.name==\"${VOLUME_NAME}\")" >/dev/null
}

exists_app || fly apps create "${APP_NAME}" --org "${ORG_NAME}"
exists_bucket || echo fly storage create --name "${BUCKET_NAME}" --org "${ORG_NAME}" --private
exists_ipv4 || fly ips allocate-v4 --app "${APP_NAME}" --yes
exists_volume || fly volumes create "${VOLUME_NAME}" --app "${APP_NAME}" --region "${REGION}" --size 1 --yes

# echo "=== Deploying app ==="
# fly deploy --local-only --strategy immediate
