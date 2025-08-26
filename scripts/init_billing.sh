#!/bin/bash
set -euo pipefail

if [[ $(uname) == Darwin ]]; then
  readonly DATE=gdate
else
  readonly DATE=date
fi

if [[ ! -d data ]]; then
  mkdir data
fi

if [[ -e data/billing.sqlite ]]; then
  sqlite3 data/billing.sqlite ".backup 'data/billing.sqlite-$(TZ=UTC ${DATE} +%Y%m%dT%H%mZ)'"
fi

while read -r file; do
  basename "${file}" >&2
  sqlite3 data/billing.sqlite <"${file}"
done < <(find db/migrations -name '*.sql' -type f | sort)
