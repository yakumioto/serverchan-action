FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v -o serverchan .

FROM alpine:3

LABEL maintainer="Mioto <yaku.mioto@gmail.com>"

RUN apk update && \
    apk add --no-cache ca-certificates

WORKDIR /drone/src

COPY --from=builder /app/serverchan /usr/local/bin/serverchan

ENTRYPOINT ["serverchan"]