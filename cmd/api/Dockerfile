ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app


COPY ../../ .
RUN go mod download
RUN go build -o ./cmd/api/bin ./cmd/api


FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*


COPY --from=builder /app/cmd/api/bin ./app/cmd/api/
COPY --from=builder /app/.env .


EXPOSE 8080 

ENTRYPOINT ./app/cmd/api/bin 
