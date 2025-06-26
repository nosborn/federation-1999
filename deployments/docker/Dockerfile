FROM debian:13.1-slim

ENV TZ=America/New_York
ENV ZONEINFO=America/New_York

COPY deployments/docker/server.crontab /tmp/server.crontab

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y update \
    && apt-get -y install --no-install-recommends \
        'awscli=2.*' \
        'busybox=1:1.*' \
        'bzip2=1.*' \
        'ca-certificates=*' \
        'inetutils-inetd=2:2.*' \
        'inetutils-telnetd=2:2.*' \
        'locales=*' \
        'opensmtpd=7.*' \
        'opensmtpd-filter-dkimsign=0.*' \
        'procps=2:4.*' \
        's6=2.*' \
        'sqlite3=3.*' \
    && rm -rf /var/lib/apt/lists/* \
    && rm -f /var/log/apt/* /var/log/dpkg.log \
    && sed -i '/^# en_US ISO-8859-1$/s/# //' /etc/locale.gen \
    && locale-gen \
    && ln -s busybox /usr/bin/crontab \
    && ln -s ../bin/busybox /usr/sbin/crond \
    && mkdir /var/spool/cron \
    && mkdir -m 1730 /var/spool/cron/crontabs \
    && rm -rf /etc/cron.* \
    && rm -f /etc/issue /etc/issue.net /etc/motd \
    && ln -sf /usr/share/zoneinfo/America/New_York /etc/localtime \
    && useradd -c 'Federation' -m -s /usr/sbin/nologin -u 1000 fed \
    && mkdir -m 0750 /home/fed/log /home/fed/log/session \
    && chown fed:fed /home/fed/log /home/fed/log/session \
    && crontab -u fed /tmp/server.crontab \
    && rm /tmp/server.crontab \
    && usermod -a -G fed www-data \
    && mkdir -m 0755 /etc/opt/fed \
    && mkdir -m 0755 /opt/fed /opt/fed/bin /opt/fed/lbin /opt/fed/lib /opt/fed/share \
    && mkdir -m 0755 /opt/ibgames /opt/ibgames/lib \
    && mkdir -m 0755 /opt/www /opt/www/sbin

ENV LANG=en_US.ISO-8859-1
ENV LC_ALL=en_US.ISO-8859-1

COPY --chmod=0500 deployments/docker/entrypoint.sh /entrypoint.sh
COPY --chmod=0400 deployments/docker/inetd.conf /etc/inetd.conf
COPY init/s6/ /etc/s6/

# COPY --chmod=0444 cmd/telnetd/issue /etc/opt/fed/issue

COPY --chmod=0555 internal/fed/lib/*.so /opt/fed/lib/
COPY --chmod=0555 internal/ibgames/lib/*.so /opt/ibgames/lib/

COPY --chmod=0555 bin/fedtpd /opt/fed/lbin/fedtpd
COPY --chmod=0555 bin/httpd /opt/www/sbin/httpd
COPY --chmod=0555 bin/login /opt/fed/bin/login
COPY --chmod=0555 bin/modemd /opt/fed/lbin/modemd
COPY --chmod=0555 bin/perivale /opt/fed/bin/perivale
COPY --chmod=0555 bin/workbench /opt/fed/bin/workbench

COPY web/public/ /var/www/htdocs/

# COPY --chmod=0640 data/person.d /person.d

WORKDIR /home/fed

COPY db/migrations/*.sql db/migrations/
COPY --chmod=0550 --chown=fed:fed start .

ENTRYPOINT ["/entrypoint.sh"]
