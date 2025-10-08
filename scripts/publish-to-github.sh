#!/bin/bash
set -euo pipefail

readonly public_remote=https://github.com/nosborn/federation-1999.git
readonly public_whitelist=(
  .clang-format
  .dockerignore
  .editorconfig
  .github/
  .gitignore
  .golangci.yaml
  .markdownlintrc
  .pre-commit-config.yaml
  COPYRIGHT.md
  Dockerfile.build
  Makefile
  README.md
  cmd/fedtpd/
  cmd/httpd/
  cmd/login/
  cmd/modemd/
  cmd/perivale/
  cmd/perivale-go/
  cmd/telnetd/
  cmd/workbench/
  data/.gitignore
  data/messages.txt
  db/migrations/
  deployments/docker/
  deployments/grafana/
  fly.yaml
  go.mod
  go.sum
  init/
  internal/collections/
  internal/config/
  internal/debug/
  internal/fed/
  internal/ibgames/
  internal/ioctl/
  internal/link/
  internal/login/
  internal/model/
  internal/monitoring/
  internal/perivale/
  internal/server/
  internal/testutil/
  internal/text/
  internal/version/
  internal/workbench/
  pkg/
  scripts/
  start
  test/
  tools/
  web/
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
