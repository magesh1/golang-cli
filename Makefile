# Variables
GO=go
MAIN_FILE=main.go
BINARY_NAME=cli-calculator

# Targets
build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_FILE)

run:
	$(GO) run $(MAIN_FILE)

clean:
	rm -f $(BINARY_NAME)

.PHONY: build run clean