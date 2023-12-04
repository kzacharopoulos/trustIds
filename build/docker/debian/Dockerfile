ARG ARCHITECTURE=amd64
FROM --platform=linux/${ARCHITECTURE} golang:1.19.8-bullseye

# Ignore this dir under a git repo
ENV GOCACHE=/home/docker/build/cache 

RUN apt-get update && apt-get install -y \
    git \
    build-essential \
    libpcap-dev \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /home/docker/build/
RUN ldconfig
ADD go.mod ./go.mod
ADD go.sum ./go.sum
RUN go mod download