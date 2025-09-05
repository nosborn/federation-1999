#!/bin/bash
set -euo pipefail

readonly public_remote=https://github.com/nosborn/federation-1999.git
readonly public_whitelist=(
  .dockerignore
  .gitattributes
  .github/
  .gitignore
  .golangci.yaml
  ARCHITECTURE-DIAGRAM.md
  COPYRIGHT.md
  Makefile
  README.md
  assets/images/telehouse25broadway.jpg
  cmd/login
  cmd/perivale
  cmd/telnetd
  data/.gitattributes
  data/.gitignore
  data/messages.txt
  db/migrations/
  deployments/docker/
  deployments/grafana/
  fly.yaml
  go.mod
  go.sum
  init/
  internal/config/
  internal/debug/
  internal/ioctl/
  internal/link/
  internal/login/
  internal/model/
  internal/modem/
  internal/monitoring/
  internal/perivale/
  internal/server/.gitattributes
  internal/server/arena/.gitattributes
  internal/server/build/
  internal/server/database/
  internal/server/global/
  internal/server/goods/
  internal/server/horsell/.gitattributes
  internal/server/jobs/
  internal/server/parser/
  internal/server/snark/.gitattributes
  internal/server/sol/.gitattributes
  internal/telnet/
  internal/text/
  internal/version/
  scripts/
  test/robobod/.keep
  tools/
)

private_repo="$(git rev-parse --show-toplevel)"
cd "${private_repo}"

private_clone=$(mktemp -d)
public_repo=$(mktemp -d)
trap 'rm -rf "${private_clone}" "${public_repo}"' EXIT
readonly private_clone public_repo

git clone --no-local "${private_repo}" "${private_clone}"
unset private_repo
git --git-dir="${private_clone}/.git" --work-tree="${private_clone}" archive HEAD -- "${public_whitelist[@]}" |
  tar -C "${public_repo}" -x
private_hash="$(git rev-parse --short HEAD)"

cd "${public_repo}"
git init
git add .
git commit -m "Initial commit [${private_hash}]"
unset private_hash

git branch -M main
git remote add origin "${public_remote}"
git push --force origin main
