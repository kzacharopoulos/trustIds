# The name of the executable
TARGET := $(shell awk '/^module/{print $$2}' go.mod)
COMMAND := $(shell dirname $$(find * -name main.go))/*.go

# Cross Compilation
#
# Usually this is easy to be done. In this case we have dependencies
# to libpcap-dev which should be natively build. This is much easier
# to do by using a proper docker image with the library natively installed
# and ids build inside it.
#
# Two helper scripts are provided:
#
# - builder.sh, that creates the proper image
# - crosscompile.sh, that runs make inside the proper image
#
# Supported platforms are:
#
# - debian (libc based)
# - alpine (musl based)
#
# Supported architectures for debian platform:
#
# - 386
# - amd64
# - arm/v5
# - arm/v7
# - arm64/v8
# - mips64le
# - ppc64le
# - s390x
#
# Supported architectures for alpine platform:
# - 386
# - amd64
# - arm/v6
# - arm/v7
# - arm64/v8
# - ppc64le
# - s390x
ARCHITECTURE := arm64/v8
PLATFORM := debian

# These will be provided to the target
VERSION := -X 'capture/internal/app.Version=$$(cat .version)'
DATE := -X 'capture/internal/app.Date=$$(date)'
BRANCH := -X 'capture/internal/app.Branch=$$(git branch --show-current)'
COMMIT := -X 'capture/internal/app.Commit=$$(git rev-parse HEAD)'
GIT_USER := -X 'capture/internal/app.GitUser=$$(git config user.name)'
GIT_EMAIL := -X 'capture/internal/app.GitEmail=$$(git config user.email)'

# Use linker flags to provide build info to the target
LDFLAGS = -ldflags "$(VERSION) $(DATE) $(BRANCH) $(COMMIT) $(GIT_USER) $(GIT_EMAIL)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")


all: fmt simplify vet build

$(TARGET): $(SRC)
	go build $(LDFLAGS) -o $(TARGET) $(COMMAND)

debug: $(SRC)
	go build -gcflags='all=-N -l' -o __debug_bin $(COMMAND)

fmt:
	@gofmt -l $(SRC)

simplify:
	@gofmt -s -d $(SRC)

vet:
	@go vet ./...

build: $(TARGET)

cross:
	@./scripts/builder.sh $(ARCHITECTURE) $(PLATFORM)
	@./scripts/crosscompile.sh $(ARCHITECTURE) $(PLATFORM) $(TARGET)

setcap: $(TARGET)
	sudo setcap 'cap_net_raw=ep' $(TARGET)

clean:
	rm -f $(TARGET)
	rm -f __debug_bin

.PHONY: fmt simplify vet build setcap clean
