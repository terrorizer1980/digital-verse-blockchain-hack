FROM golang:1.16-alpine

ENV DV_PORT 8080
ENV APP_NAME DigitalVerse

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh linux-headers
RUN apk add build-base

ADD ./  /go/src/${APP_NAME}/

WORKDIR /go/src/${APP_NAME}/

RUN go mod download

RUN go build -o /go/bin/${APP_NAME} .

ENTRYPOINT /go/bin/${APP_NAME}

EXPOSE ${DV_PORT}