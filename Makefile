GO ?= go

.PHONY: build
build:
	$(GO) build -o bin/app & cp .ini /bin/.ini
.PHONY: run
run:
	$(GO) run main.go

.PHONY: start
start:
	$(GO) build -o bin/app && ./app >> app.log 2>&1

.PHONY: generate
generate:
	wire .
