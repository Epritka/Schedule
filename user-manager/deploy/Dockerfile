ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update \
    && apk add --no-cache \
    make=4.3-r0
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

FROM alpine:3.13

RUN apk update
WORKDIR /go/src/app

RUN adduser -D appuser && chown -R appuser /go/src/app
COPY --from=builder /go/src/app/deploy/bin ./bin
COPY --from=builder /go/src/app/scripts ./scripts
COPY --from=builder /go/src/app/migrations ./migrations
RUN chmod +x -R ./bin ./scripts
USER appuser

CMD ["./bin/http"]
