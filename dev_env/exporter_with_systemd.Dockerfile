

FROM postgres:16-bullseye

RUN apt-get -qq -y update
RUN apt-get -qq -y install barman cron nano
RUN apt-get -qq -y install --no-install-recommends systemd systemd-sysv dbus dbus-user-session

COPY --from=exporter:latest /app/exporter /opt/barman_exporter
RUN chmod a+x /opt/barman_exporter

USER barman

# Setup cron jobs for barman, every minute run barman cron
RUN echo "* * * * * barman cron" >> ~/cron_jobs
RUN crontab ~/cron_jobs

COPY barman.conf /etc/barman.d/pg.conf
COPY barman_setup/* /opt

COPY exporter_setup/config.yaml /etc/barman_exporter/config.yaml

USER root

ENV container docker
STOPSIGNAL SIGRTMIN+3
VOLUME [ "/tmp", "/run", "/run/lock" ]
WORKDIR /

RUN rm -f /lib/systemd/system/multi-user.target.wants/* \
  /etc/systemd/system/*.wants/* \
  /lib/systemd/system/local-fs.target.wants/* \
  /lib/systemd/system/sockets.target.wants/*udev* \
  /lib/systemd/system/sockets.target.wants/*initctl* \
  /lib/systemd/system/sysinit.target.wants/systemd-tmpfiles-setup* \
  /lib/systemd/system/systemd-update-utmp*

COPY sysd_unit/exporter.service /etc/systemd/system/exporter.service

CMD [ "/lib/systemd/systemd", "log-level=info", "unit=sysinit.target" ]

