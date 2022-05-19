FROM golang

ENV APP_DIR=/app

WORKDIR /src

COPY *.go go.mod .

RUN mkdir -p $APP_DIR/tmp && \
    go build -o $APP_DIR/app && \
    useradd --home-dir $APP_DIR --shell /bin/false app && \
    chown -R app:app $APP_DIR

WORKDIR $APP_DIR

# Unmounted, this will create a new ephemeral volume on the host using the local driver
VOLUME /app/tmp

ENTRYPOINT [ "/app/app" ]
