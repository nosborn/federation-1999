FROM alpine:3.22.1

ARG ARCH=amd64

COPY deployments/docker/server.crontab /tmp/server.crontab

RUN apk add --no-cache aws-cli=2.27.25-r0 sqlite=3.49.2-r1 supervisor=4.2.5-r5 tzdata=2025b-r0 util-linux=2.41-r9 \
    && apk add --no-cache curl=8.14.1-r1 \
    && rm /etc/crontabs/* \
    && ln -sf /usr/share/zoneinfo/EST5EDT /etc/localtime \
    && adduser -s /sbin/nologin -D -u 1000 fed \
    && mkdir -m 0750 ~fed/log ~fed/log/session \
    && chown fed:fed ~fed/log ~fed/log/session \
    && crontab -u fed /tmp/server.crontab \
    && rm /tmp/server.crontab \
    && mkdir -m 0755 /opt/fed /opt/fed/bin /opt/fed/lbin /opt/fed/share

COPY --chmod=0500 deployments/docker/entrypoint.sh /entrypoint.sh
COPY --chmod=0400 init/supervisord.conf /etc/supervisord.conf

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
