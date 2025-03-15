.PHONY: all build run run-interactive run-small run-medium run-large clean download-words help

# Default target
all: download-words build

# Build the Go application
build:
	@echo "Building document generator..."
	@go build -o document-generator main.go

# Run the application with default settings
run: check-words build
	@echo "Running document generator with default settings..."
	@./document-generator

# Run the application in interactive mode
run-interactive: check-words build
	@echo "Running document generator in interactive mode..."
	@./document-generator -interactive

# Run with small dataset (100 documents, shorter length)
run-small: check-words build
	@echo "Running document generator with small dataset..."
	@./document-generator -num 100 -min 5 -max 50 -avg 20 -output documents-small.txt

# Run with medium dataset (10,000 documents)
run-medium: check-words build
	@echo "Running document generator with medium dataset..."
	@./document-generator -num 10000 -min 10 -max 500 -avg 100 -output documents-medium.txt

# Run with large dataset (default 1,000,000 documents)
run-large: check-words build
	@echo "Running document generator with large dataset..."
	@./document-generator

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f document-generator
	@rm -f documents*.txt

# Download English words from GitHub
download-words:
	@echo "Downloading English words list..."
	@curl -s https://raw.githubusercontent.com/first20hours/google-10000-english/master/google-10000-english.txt > words.txt
	@echo "Downloaded $(shell wc -l < words.txt) English words"

# Check if the words file exists and has content
check-words:
	@if [ ! -s words.txt ]; then \
		echo "Words file is empty or does not exist. Downloading..."; \
		$(MAKE) download-words; \
	else \
		echo "Words file exists with $(shell wc -l < words.txt) words"; \
	fi

# Help command to show available options
help:
	@echo "Document Generator - Available commands:"
	@echo "  make build            - Build the application"
	@echo "  make run              - Run with default settings (1M documents)"
	@echo "  make run-interactive  - Run in interactive mode (prompts for settings)"
	@echo "  make run-small        - Run with small dataset (100 documents)"
	@echo "  make run-medium       - Run with medium dataset (10K documents)"
	@echo "  make run-large        - Run with large dataset (1M documents)"
	@echo "  make clean            - Clean build artifacts"
	@echo "  make download-words   - Download English words list"
	@echo "  make help             - Show this help message"
	@echo ""
	@echo "Command-line flags (when running directly):"
	@echo "  -num N         - Number of documents to generate"
	@echo "  -min N         - Minimum words per document"
	@echo "  -max N         - Maximum words per document"
	@echo "  -avg N         - Target average words per document"
	@echo "  -output FILE   - Output file path"
	@echo "  -words FILE    - English words file path"
	@echo "  -interactive   - Run in interactive mode"
