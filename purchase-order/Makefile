# https://ops.tips/blog/minimal-golang-makefile/

VERSION         :=      $(shell cat ./VERSION)

all: install

install:
	go install -v

test:
	go test ./... -v

fmt:
	go fmt ./... -v

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

.PHONY: install test fmt release
