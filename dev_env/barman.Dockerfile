

FROM postgres:16-bullseye

RUN apt-get -qq -y update
RUN apt-get -qq -y install barman cron nano


COPY --from=exporter:latest /app/exporter /opt/exporter
RUN chmod a+x /opt/exporter

USER barman

# Setup cron jobs for barman, every minute run barman cron
RUN echo "* * * * * barman cron" >> ~/cron_jobs
RUN crontab ~/cron_jobs

COPY barman.conf /etc/barman.d/pg.conf
COPY barman_setup/* /opt

CMD ["/opt/start_barman.sh"]