FROM debian:12.11-slim

ARG ARCH=amd64

COPY deployments/docker/server.crontab /tmp/server.crontab

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y update \
    && apt-get -y install --no-install-recommends \
        'apache2=2.4.*' \
        'awscli=2.9.*' \
        'bzip2=1.0.*' \
        'cron=3.*' \
        'libapache2-mod-perl2=2.0.*' \
        'sqlite3=3.40.*' \
        'supervisor=4.2.*' \
    && rm -rf /var/lib/apt/lists/* \
    && rm /etc/cron.d/e2scrub_all \
    && rm /etc/cron.daily/apt-compat /etc/cron.daily/dpkg \
    && ln -sf /usr/share/zoneinfo/EST5EDT /etc/localtime \
    && useradd -c 'Federation' -m -s /usr/sbin/nologin -u 1000 fed \
    && mkdir -m 0750 /home/fed/log /home/fed/log/session \
    && chown fed:fed /home/fed/log /home/fed/log/session \
    && crontab -u fed /tmp/server.crontab \
    && rm /tmp/server.crontab \
    && mkdir -m 0755 /opt/fed /opt/fed/bin /opt/fed/lbin /opt/fed/share

COPY --chmod=0500 deployments/docker/apache2.conf /etc/apache2/apache2.conf
COPY --chmod=0500 deployments/docker/entrypoint.sh /entrypoint.sh
COPY --chmod=0400 init/supervisord.conf /etc/supervisor/supervisord.conf

COPY --chmod=0555 bin/linux-${ARCH}/fedtpd /opt/fed/lbin/fedtpd
COPY --chmod=0555 bin/linux-${ARCH}/login /opt/fed/bin/login
COPY --chmod=0555 bin/linux-${ARCH}/perivale /opt/fed/bin/perivale
COPY --chmod=0555 bin/linux-${ARCH}/telnetd /opt/fed/lbin/telnetd
COPY --chmod=0555 bin/linux-${ARCH}/workbench /opt/fed/bin/workbench

# COPY --chmod=0640 data/person.d /person.d

WORKDIR /home/fed

COPY db/migrations/*.sql db/migrations/
COPY --chmod=0550 --chown=fed:fed start .

ENTRYPOINT ["/entrypoint.sh"]
