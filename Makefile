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

combined_test:
	go run ./cmd/karas/karas.go --json ./misc/data/karas.json \
		--chrome "./drivers/chrome-linux/chrome" \
		--chrome-driver "./drivers/chromedriver" \
		--selenium "./drivers/selenium-server.jar"
combined_test_for_config:
	go run ./cmd/karas/karas.go --config ./misc/data/Karasfile \
		--chrome "./drivers/chrome-linux/chrome" \
		--chrome-driver "./drivers/chromedriver" \
		--selenium "./drivers/selenium-server.jar"
