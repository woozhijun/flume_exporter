FROM golang:1.11
MAINTAINER woozhijun


EXPOSE 9360
ENTRYPOINT [ "/bin/flume_exporter" ]
