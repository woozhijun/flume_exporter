FROM       quay.io/prometheus/busybox:latest
MAINTAINER woozhijun

COPY flume_exporter /bin/flume_exporter

RUN mkdir -p /etc/flume_exporter
COPY metrics.yml /etc/flume_exporter/metrics.yml
COPY config.yml /etc/flume_exporter/config.yml

EXPOSE 9360
ENTRYPOINT [ "/bin/flume_exporter", "--metric-file=/etc/flume_exporter/metrics.yml", "--config-file=/etc/flume_exporter/config.yml" ]
