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

The project includes a Makefile to simplify common tasks:

```bash
# Download the English words file (required before first run)
make download-words

# Build the application
make build

# Run the document generator
make run

# Clean up build artifacts
make clean
```

Alternatively, you can run the program directly:

```bash
go run main.go
```

The program will create a file named `documents.txt` in the current directory.

## Configuration

You can modify the following constants in the code to adjust the output:

- `numDocuments`: Total number of documents to generate (default: 1,000,000)
- `minWordsPerDoc`: Minimum words per document (default: 10)
- `maxWordsPerDoc`: Maximum words per document (default: 1,000)
- `avgWordsPerDoc`: Target average words per document (default: 200)
- `wordsFile`: Path to the file containing English words (default: "words.txt")

## Performance

The program uses buffered I/O for better performance when writing the large output file. It loads a dictionary of real English words from a file and reuses them for all documents.

The triangular distribution is used to achieve the target average document length while maintaining the specified minimum and maximum document lengths.
