

.PHONY: all
all: test repl

.PHONY: repl
repl:
	go build -o bin/repl ./cmd/repl/*

.PHONY: test
test:
	go test ./...

