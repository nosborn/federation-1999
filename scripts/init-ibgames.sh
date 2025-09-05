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

if [[ -e data/accounts.sqlite ]]; then
  sqlite3 data/accounts.sqlite ".backup 'data/accounts.sqlite-$(TZ=UTC ${DATE} +%Y%m%dT%H%mZ)'"
fi

while read -r file; do
  basename "${file}" >&2
  sqlite3 data/accounts.sqlite <"${file}"
done < <(find db/migrations -name '*.sql' -type f | sort)
