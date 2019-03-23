NAME := geso
VERSION := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X github.com/bsoo/geso/cmd.Version=${VERSION} -X github.com/bsoo/geso/cmd.Revision=${REVISION}"

.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -ldflags $(LDFLAGS) -o bin/$(NAME)

.PHONY: install
install:
	go install -ldflags $(LDFLAGS)
