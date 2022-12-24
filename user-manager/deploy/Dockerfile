ARG GO_VERSION=1.19
ARG ALPINE_VERSION=3.17

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

RUN apk update \
    && apk add --no-cache \
    make
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

FROM alpine:${ALPINE_VERSION}

RUN apk update
WORKDIR /go/src/app

RUN adduser -D appuser && chown -R appuser /go/src/app
COPY --from=builder /go/src/app/deploy/bin ./bin
COPY --from=builder /go/src/app/scripts ./scripts
COPY --from=builder /go/src/app/migrations ./migrations
RUN chmod +x -R ./bin ./scripts
USER appuser

CMD ["./bin/server"]
