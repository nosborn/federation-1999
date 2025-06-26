#!/bin/bash
set -euo pipefail

if [[ $(uname) == Darwin ]]; then
  readonly DATE=gdate
else
  readonly DATE=date
fi
if [[ -e "${DBPATH:?}/ibgames.sqlite" ]]; then
  sqlite3 "${DBPATH:?}/ibgames.sqlite" ".backup '${DBPATH:?}/ibgames.sqlite-$(TZ=UTC ${DATE} +%Y%m%dT%H%mZ)'"
fi

while read -r file; do
  case "${file}" in
    *.gz) CAT=gzcat ;;
  esac
  basename "${file}" >&2
  "${CAT:-cat}" "${file}" | sqlite3 "${DBPATH:?}/ibgames.sqlite"
  unset CAT
done < <(find db/migrations -name '*.sql' -o -name '*.sql.gz' | sort)
