.PHONY: all build run clean download-words

# Default target
all: download-words build

# Build the Go application
build:
	@echo "Building document generator..."
	@go build -o document-generator main.go

# Run the application
run: build
	@echo "Running document generator..."
	@./document-generator

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f document-generator
	@rm -f documents.txt

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

# Generate documents with English words
generate: check-words build
	@echo "Generating documents with English words..."
	@./document-generator
