#!/bin/sh
# set -euo pipefail

# timestamp="$(TZ=UTC date +%Y%m%dT%H%M%SZ)"
# readonly timestamp

if [ -n "${FLY_REGION:-}" ]; then
  _hostname="$(hostname)"
  if [ "${_hostname}" != federation-1999.fly.dev ]; then
    hostname federation-1999.fly.dev
    sed -i "s/${_hostname}/federation-1999.fly.dev\tfederation-1999/" /etc/hosts
  fi
  unset _hostname
fi

exec /usr/bin/s6-svscan /etc/s6

# for dir in backup lock log log/session run; do
#   [[ -d "/home/fed/${dir}" ]] || mkdir "/home/fed/${dir}"
#   chmod 0750 "/home/fed/${dir}"
#   chown fed:fed "/home/fed/${dir}"
# done
#
# if [[ -n "${FLY_REGION:-}" ]]; then
#   chmod 0750 /home/fed/data
#   chown fed:fed /home/fed/data
#
#   for n in $(seq 0 9); do
#     mkdir -p "/home/fed/data/workbench${n}"
#     chmod 0750 "/home/fed/data/workbench${n}"
#     chown fed:fed "/home/fed/data/workbench${n}"
#   done
# fi
#
# # if [ ! -e data/ibgames.sqlite ] && [ -n "${BUCKET_NAME:-}" ]; then
# #   : # TODO
# # fi
# if [[ -n "${FLY_REGION:-}" ]]; then
#   chown fed:fed data/ibgames.sqlite
# fi
# # if [ -e data/ibgames.sqlite ]; then
# #   _backup="data/ibgames.sqlite-${timestamp}"
# #   sqlite3 -cmd '.bail on' -readonly data/ibgames.sqlite ".backup '${_backup}'"
# #   bzip2 -9 "${_backup}"
# #   if [ -n "${BUCKET_NAME:-}" ]; then
# #     aws s3 cp "${_backup}.bz2" "s3://${BUCKET_NAME}/${timestamp}/ibgames.sqlite.bz2" --no-progress
# #   fi
# #   unset _backup
# # fi
# # find data -maxdepth 1 -mtime +7 -name 'ibgames.sqlite-*.bz2' -type f -delete
# # if [ -d db/migrations ]; then
# #   find db/migrations -name '*.sql' -type f | sort | while read -r file; do
# #     basename "${file}" >&2
# #     sqlite3 data/ibgames.sqlite <"${file}"
# #   done
# # fi
# #
# # if [ ! -f data/person.d ] && [ -n "${BUCKET_NAME:-}" ]; then
# #   : # TODO
# # fi
# # # if [ ! -f data/person.d ]; then
# # #   tar -xjf /data-*.tar.bz2
# # #   chown -R fed:fed data
# # #   find data -type d -exec chmod 0770 {} \;
# # #   find data -type f -exec chmod 0660 {} \;
# # # fi
# # if [ -e data/person.d ]; then
# #   _backup="data/person.d-${timestamp}"
# #   cp -pv data/person.d "${_backup}"
# #   bzip2 -9 "${_backup}"
# #   if [ -n "${BUCKET_NAME:-}" ]; then
# #     aws s3 cp "${_backup}.bz2" "s3://${BUCKET_NAME}/${timestamp}/person.d.bz2" --no-progress
# #   fi
# #   unset _backup
# # fi
# # find data -maxdepth 1 -mtime +7 -name 'person.d-*.bz2' -type f -delete
# #
# # for n in $(seq 0 9); do
# #   _dir="data/workbench${n}"
# #   if [ ! -d "${_dir}" ]; then
# #     mkdir -m 0750 "${_dir}"
# #   fi
# #   unset _dir
# # done
# # _backup="data/workbench-${timestamp}"
# # env BZIP2=--best tar -cjf "${_backup}.tar.bz2" data/workbench[0-9]/
# # if [ -n "${BUCKET_NAME:-}" ]; then
# #   aws s3 cp "${_backup}.tar.bz2" "s3://${BUCKET_NAME}/${timestamp}/workbench.tar.bz2" --no-progress
# # fi
# # unset _backup
# # find data -maxdepth 1 -mtime +7 -name 'workbench-*.tar.bz2' -type f -delete
#
# if [[ -n "${FLY_REGION:-}" ]]; then
#   export FEDTPD_FLAGS='-check -free-period -test-features -trace' # TODO: '-free-period'
#   export HTTPD_FLAGS='-proxy-proto'
# else
#   export FEDTPD_FLAGS='-check -free-period -test-features -trace'
#   export HTTPD_FLAGS=
# fi
# exec /usr/bin/supervisord -c /etc/supervisor/supervisord.conf -n
