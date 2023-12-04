#!/usr/bin/env bash
#

readonly ARCHITECTURE="$1"
readonly PLATFORM="$2"
readonly DOCKER_PATH="build/docker/${PLATFORM}"
readonly IMAGE_TAG="ids-builder-$(echo "$ARCHITECTURE" | tr "/" "-")-${PLATFORM}:latest"

go mod tidy

cp go.mod "$DOCKER_PATH"
cp go.sum "$DOCKER_PATH"

pushd "$DOCKER_PATH"
DOCKER_BUILDKIT=1 docker build -q \
	--build-arg ARCHITECTURE="$ARCHITECTURE" \
	--build-arg USER_ID="$(id -u "$USER")" \
	--build-arg GROUP_ID="$(id -g "$USER")" \
	--platform "linux/${ARCHITECTURE}" \
	-t "$IMAGE_TAG" .
popd

rm "$DOCKER_PATH/go.mod"
rm "$DOCKER_PATH/go.sum"
