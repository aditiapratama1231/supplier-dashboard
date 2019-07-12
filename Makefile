# https://ops.tips/blog/minimal-golang-makefile/

VERSION         :=      $(shell cat ./VERSION)

all: install

install:
	go install -v

test:
	go test ./... -v

fmt:
	go fmt ./... -v

migrate:
	@which soda > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u -v github.com/gobuffalo/pop/...; \
		go install -v github.com/gobuffalo/pop/soda; \
	fi
	soda migrate -p database up

migrate_reset:
	@which soda > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u -v github.com/gobuffalo/pop/...; \
		go install -v github.com/gobuffalo/pop/soda; \
	fi
	soda migrate -p database reset

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

.PHONY: install test fmt release
