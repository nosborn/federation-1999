FROM debian:13.1-slim

COPY deployments/docker/server.crontab /tmp/server.crontab

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y update \
    && apt-get -y install --no-install-recommends \
        'awscli=2.23.*' \
        'bzip2=1.0.*' \
        'ca-certificates=*' \
        'cron=3.*' \
        'opensmtpd=7.6.*' \
        'opensmtpd-filter-dkimsign=0.*' \
        'procps=*' \
        'sqlite3=3.46.*' \
        'supervisor=4.2.*' \
    && rm -rf /var/lib/apt/lists/*

RUN rm /etc/cron.daily/apt-compat /etc/cron.daily/dpkg \
    && ln -sf /usr/share/zoneinfo/America/New_York /etc/localtime \
    && useradd -c 'Federation' -m -s /usr/sbin/nologin -u 1000 fed \
    && mkdir -m 0750 /home/fed/log /home/fed/log/session \
    && chown fed:fed /home/fed/log /home/fed/log/session \
    && crontab -u fed /tmp/server.crontab \
    && rm /tmp/server.crontab \
    && mkdir -m 0755 /opt/fed /opt/fed/bin /opt/fed/lbin /opt/fed/share \
    && mkdir -m 0755 /opt/www /opt/www/sbin

COPY --chmod=0500 deployments/docker/entrypoint.sh /entrypoint.sh
COPY --chmod=0400 init/supervisord.conf /etc/supervisor/supervisord.conf

COPY --chmod=0555 bin/linux-amd64/fedtpd /opt/fed/lbin/fedtpd
COPY --chmod=0555 bin/linux-amd64/httpd /opt/www/sbin/httpd
COPY --chmod=0555 bin/linux-amd64/login /opt/fed/bin/login
COPY --chmod=0555 bin/linux-amd64/perivale /opt/fed/bin/perivale
COPY --chmod=0555 bin/linux-amd64/telnetd /opt/fed/lbin/telnetd
COPY --chmod=0555 bin/linux-amd64/workbench /opt/fed/bin/workbench

COPY web/public/ /var/www/htdocs/

# COPY --chmod=0640 data/person.d /person.d

WORKDIR /home/fed

COPY db/migrations/*.sql db/migrations/
COPY --chmod=0550 --chown=fed:fed start .

ENTRYPOINT ["/entrypoint.sh"]
