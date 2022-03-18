ARG APP_NAME=echo

FROM golang:1.18 AS builder
ARG APP_NAME

# Setup Git:
# Authorize SSH Host
RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

RUN mkdir -p /go/src/github.com/ricardo-ch/$APP_NAME
ADD . /go/src/github.com/ricardo-ch/$APP_NAME
WORKDIR /go/src/github.com/ricardo-ch/$APP_NAME

RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app -a -ldflags '-s' -installsuffix cgo main.go

# Forget everything, just get the app binary
FROM debian:buster-slim
ARG APP_NAME

RUN apt-get update && apt-get install -y ca-certificates

COPY --chown=1984:1984 --from=builder /go/src/github.com/ricardo-ch/$APP_NAME/app /

EXPOSE 1323

# Non-root user
USER 1984:1984

CMD ["/app"]