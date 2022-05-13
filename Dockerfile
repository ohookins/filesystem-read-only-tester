FROM golang

WORKDIR /src

COPY *.go go.mod .

RUN mkdir -p /app/tmp && \
    go build -o /app/app

WORKDIR /app

# TODO: Uncomment later and see if it makes a difference
# VOLUME /app/tmp

ENTRYPOINT [ "/app/app" ]
