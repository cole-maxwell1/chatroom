# Simple Makefile for a Go project
BINARY_NAME=chatroom

# Build the application
all: build

templ:
	@if command -v templ > /dev/null; then \
	    templ generate; \
	    echo "templ templates generated!";\
	else \
	    read -p "`templ` is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/a-h/templ/cmd/templ@latest; \
	        templ generate; \
	        echo "templ templates generated!";\
	    else \
	        echo "You chose not to install templ. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

tailwind:
	@if command -v npm > /dev/null; then \
	    npm run tailwind:build; \
	    echo "css built!";\
	else \
	    echo "npm is not installed on your machine."; \
		echo "Please install npm and nodejs on you machine."; \
		echo "Exiting..."; \
		exit 1; \
	fi


build: templ tailwind
	@echo "Compinling go binary..."
	@go build -v -o ./tmp/${BINARY_NAME} cmd/${BINARY_NAME}/main.go

# Run the application
run: templ tailwind
	@go run cmd/${BINARY_NAME}/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f tmp/${BINARY_NAME}

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
