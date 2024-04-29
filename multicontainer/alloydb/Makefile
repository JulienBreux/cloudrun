NAME    ?= mksctl
BINDIR  := $(CURDIR)/bin
COMMIT  ?= $(shell git rev-parse HEAD)
DATE    ?= $(shell date +%FT%T%Z)

build: ## [binary] Build local binary
	mkdir -p ./bin
	CGO_ENABLED=0 go build -ldflags='-s -w' -o ${BINDIR}/app .
