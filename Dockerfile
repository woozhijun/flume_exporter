FROM       quay.io/prometheus/busybox:latest
MAINTAINER woozhijun

COPY flume_exporter /bin/flume_exporter

EXPOSE 9360
ENTRYPOINT [ "/bin/flume_exporter" ]