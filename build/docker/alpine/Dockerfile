ARG ARCHITECTURE=amd64
FROM --platform=linux/${ARCHITECTURE} golang:1.19.8-alpine3.17

# Ignore this dir under a git repo
ENV GOCACHE=/home/docker/build/cache 

RUN apk add --no-cache git build-base musl-dev libpcap-dev
WORKDIR /home/docker/build/
ADD go.mod ./go.mod
ADD go.sum ./go.sum
RUN git config --global --add safe.directory /home/docker/build
RUN go mod download
