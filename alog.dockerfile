FROM golang:alpine

RUN export CONFIG_FILE=./configs/config.yaml

COPY bin/alog_test /usr/local/bin/

ENTRYPOINT [ "alog_test", "run" ]
