.PHONY: all build run run-interactive run-small run-medium run-large clean download-words check-words help

# Default target
all: check-words build

# Build the application
build:
	go build -o document-generator cmd/main.go

# Run with default settings (1M documents)
run: check-words
	./document-generator

# Run in interactive mode
run-interactive: check-words
	./document-generator -interactive

# Run with small dataset (100 documents)
run-small: check-words
	./document-generator -num 100 -min 5 -max 50 -avg 20

# Run with medium dataset (10K documents)
run-medium: check-words
	./document-generator -num 10000 -min 10 -max 500 -avg 100

# Run with large dataset (1M documents)
run-large: check-words
	./document-generator -num 1000000

# Clean build artifacts
clean:
	rm -f document-generator documents.txt

# Download English words list
download-words:
	@echo "Downloading English words list..."
	@curl -s -o words.txt https://raw.githubusercontent.com/dwyl/english-words/master/words.txt
	@echo "Words file downloaded successfully"

# Check if words file exists and has content
check-words:
	@if [ ! -s words.txt ]; then \
		echo "Words file missing or empty. Downloading..."; \
		$(MAKE) download-words; \
	fi

# Show help
help:
	@echo "Available commands:"
	@echo "  make all              - Download words and build"
	@echo "  make build            - Build the application"
	@echo "  make run              - Run with default settings (1M documents)"
	@echo "  make run-interactive  - Run in interactive mode"
	@echo "  make run-small        - Run with small dataset (100 documents)"
	@echo "  make run-medium       - Run with medium dataset (10K documents)"
	@echo "  make run-large        - Run with large dataset (1M documents)"
	@echo "  make clean            - Remove build artifacts"
	@echo "  make download-words   - Download English words list"
	@echo "  make check-words      - Check if words file exists"
	@echo "  make help             - Show this help message"
