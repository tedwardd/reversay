BIN := bin/echo-reverse

.PHONY: build run

build:
	go build -o $(BIN) .

run: build
	./$(BIN)
