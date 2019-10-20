FROM golang:1.13.3-alpine3.10

WORKDIR /go/src/app

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/oxequa/realize

EXPOSE 8080

CMD ["realize", "start"]
