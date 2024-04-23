.PHONY: build install

VERSION := $(or $(VERSION),"0.0.0-$(shell git rev-parse --short HEAD)")

build:
	go build -trimpath -ldflags "-s -w -X main.version=$(VERSION)"

install:
	go install -trimpath -ldflags "-s -w -X main.version=$(VERSION)"
