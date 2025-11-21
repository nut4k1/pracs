FROM golang:1.25-alpine

ARG PACKAGES="vim openssl-dev build-base curl less git bash screen"

RUN apk update && apk add --no-cache $PACKAGES

WORKDIR /app
COPY . /app