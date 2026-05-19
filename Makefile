B = $(shell git rev-parse --abbrev-ref HEAD)
BRANCH = $(subst /,-,$(B))
GITREV = $(shell git describe --abbrev=7 --always --tags --dirty)
FULLREV = $(GITREV)-$(BRANCH)-$(shell date +%Y%m%d-%H:%M:%S)
VERSION = $(shell cat VERSION)

BIN_NAME = kvs

LDFLAGS = "-X main.fullrev=$(FULLREV) -X main.version=$(VERSION) -s -w"

all: native

mk_build_dir:
	mkdir -p build

native:
	go build -ldflags=$(LDFLAGS) -o build/$(BIN_NAME)

linux_amd64: mk_build_dir
	GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/$(BIN_NAME)

linux_arm: mk_build_dir
	CGO_ENABLED=0 GOARM=5 GOOS=linux GOARCH=arm go build -ldflags=$(LDFLAGS) -o build/$(BIN_NAME)arm

win: mk_build_dir
	GOOS=windows GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o build/$(BIN_NAME).exe

install:
	cp -v build/$(BIN_NAME) /usr/local/bin

uninstall:
	rm /usr/local/bin/$(BIN_NAME)

clean:
	rm -rf build
