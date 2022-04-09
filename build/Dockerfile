FROM alpine:latest

ENV ROOT_DIR="/app"

RUN mkdir -p $ROOT_DIR

COPY ./greet $ROOT_DIR/greet
COPY ./etc/greet-api.yaml $ROOT_DIR/etc/greet-api.yaml

WORKDIR $ROOT_DIR

ENTRYPOINT ["./greet"]