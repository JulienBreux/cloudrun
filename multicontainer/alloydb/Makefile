NAME    ?= todos
BINDIR  := $(CURDIR)/bin

build: ## [binary] Build local binary
	mkdir -p ./bin
	CGO_ENABLED=0 go build -ldflags='-s -w' -o ${BINDIR}/app .
