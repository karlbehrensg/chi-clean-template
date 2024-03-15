# Run the application
run:
	@go run cmd/api/main.go

# Build the application
build:
	@echo "Building..."
	@go build -o bin/main cmd/api/main.go

# Start the application
start:
	@echo "Starting..."
	@./bin/main

# Live Reload
dev:
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

.PHONY: run build start dev
