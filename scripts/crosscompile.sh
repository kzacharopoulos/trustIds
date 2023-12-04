#!/usr/bin/env bash
#

readonly ARCHITECTURE="$1"
readonly PLATFORM="$2"
readonly TARGET="$3"
readonly IMAGE_TAG="ids-builder-$(echo "$ARCHITECTURE" | tr "/" "-")-${PLATFORM}:latest"

docker run --rm \
	-u "$(id -u "$USER"):$(id -g "$USER")" \
	--platform="linux/${ARCHITECTURE}" \
	-v "$PWD":/home/docker/build \
	"$IMAGE_TAG" make "$TARGET"
