.PHONY: build clean install uninstall

APP_NAME := karas
PREFIX := /usr/local/bin

install:
	./scripts/install.sh --prefix $(PREFIX)
uninstall:
	./scripts/uninstall.sh --prefix $(PREFIX)

build:
	./scripts/build.sh --clean

clean:
	rm -fr ./bin/*

test:
	go test -v ./...

lint:
	golangci-lint run
