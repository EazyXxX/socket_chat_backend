.SILENT:

BINARY_NAME=server

build:
		go build -o $(BINARY_NAME)

clean:
		rm -f $(BINARY_NAME)

dev: build
		./$(BINARY_NAME)

.DEFAULT_GOAL := dev
