# Document Generator

A Go program that generates a large text file (`documents.txt`) containing documents with real English words.

## Features

- Generates 1 million documents in a single text file
- Each document is on its own line
- Documents contain between 10-1000 real English words
- Average document length is 200 words
- Uses a dictionary of real English words from Google's most common English words
- Shows progress during generation

## Usage

The project includes a Makefile with various commands to simplify usage:

```bash
# Default target - download words and build
make all

# Build the application
make build

# Run with default settings (1M documents)
make run

# Run in interactive mode (prompts for settings)
make run-interactive

# Run with small dataset (100 documents, shorter length)
make run-small

# Run with medium dataset (10K documents)
make run-medium

# Run with large dataset (1M documents)
make run-large

# Clean build artifacts (removes executable and documents files)
make clean

# Download English words list from GitHub
make download-words

# Check if the words file exists and has content
make check-words

# Show help with all available commands
make help
```

### Makefile Command Details

- `make all`: Default target that downloads words and builds the application
- `make build`: Compiles the Go application into an executable named "document-generator"
- `make run`: Runs the application with default settings (1M documents)
- `make run-interactive`: Runs the application in interactive mode, prompting for settings
- `make run-small`: Runs with a small dataset (100 documents, 5-50 words per doc, avg 20 words)
- `make run-medium`: Runs with a medium dataset (10K documents, 10-500 words per doc, avg 100 words)
- `make run-large`: Runs with a large dataset (default 1M documents)
- `make clean`: Removes the executable and all generated documents files
- `make download-words`: Downloads a list of English words from GitHub
- `make check-words`: Verifies the words file exists and has content, downloads if needed
- `make help`: Displays help information about available commands

### Command-line Flags

You can also run the program directly with command-line flags to customize the generation:

```bash
# Run with custom settings
./document-generator -num 5000 -min 20 -max 300 -avg 100 -output custom-docs.txt

# Run in interactive mode
./document-generator -interactive
```

Available flags:

- `-num N`: Number of documents to generate (default: 1,000,000)
- `-min N`: Minimum words per document (default: 10)
- `-max N`: Maximum words per document (default: 1,000)
- `-avg N`: Target average words per document (default: 200)
- `-output FILE`: Output file path (default: "documents.txt")
- `-words FILE`: English words file path (default: "words.txt")
- `-interactive`: Run in interactive mode

### Interactive Mode

When running in interactive mode, the program will prompt you for most configuration values, showing the current/default value in brackets. Press Enter to keep the current value, or type a new value to change it.

The program will prompt for:
- Number of documents to generate
- Minimum words per document
- Maximum words per document
- Target average words per document
- Output file path

Note: The words file path is not prompted for in interactive mode. It will use either the default value ("words.txt") or the value provided via the command-line flag (`-words`).

#### Input Handling

The interactive mode features improved input handling:
- Default values are clearly displayed in brackets [like this]
- Simply pressing Enter will use the default value
- Input is properly trimmed of whitespace and newline characters
- Error handling ensures the program continues even if input reading fails

## Configuration

The document generator can be configured using command-line flags or interactive mode, without needing to modify the code. The following settings can be adjusted:

- Number of documents to generate (default: 1,000,000)
- Minimum words per document (default: 10)
- Maximum words per document (default: 1,000)
- Target average words per document (default: 200)
- Output file path (default: "documents.txt")
- English words file path (default: "words.txt")

See the [Usage](#usage) section for details on how to specify these settings.

## Performance

The program uses buffered I/O for better performance when writing the large output file. It loads a dictionary of real English words from a file and reuses them for all documents.

The triangular distribution is used to achieve the target average document length while maintaining the specified minimum and maximum document lengths.
