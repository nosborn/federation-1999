#!/bin/bash
set -euo pipefail

BIN_DIR="${1:-}"
if [[ -z "${BIN_DIR}" ]]; then
  BIN_DIR="bin/$(go env GOOS)-$(go env GOARCH)"
fi

export LANG=en_US.iso88591

trap '[[ -n "${telnetd:-}" ]] && kill "${telnetd}"' EXIT INT TERM
"./${BIN_DIR}/telnetd" </dev/null >/dev/null &
telnetd=$!

mkdir -p lock
mkdir -p log/session
mkdir -p run

"./${BIN_DIR}/fedtpd" -check -free-period -test-features -trace </dev/null 2>&1 | tee log/dev.log
