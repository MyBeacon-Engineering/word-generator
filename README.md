# Document Generator

A Go program that generates text files containing random documents composed of real English words. The program uses a triangular distribution to control document lengths, ensuring a natural variation while maintaining a target average length.

## Features

- Generates a specified number of documents
- Uses real English words from a dictionary
- Configurable document length parameters:
  - Minimum words per document
  - Maximum words per document
  - Target average words per document
- Progress tracking during generation
- Interactive mode for configuration
- Command-line flags for all parameters

## Installation

1. Clone the repository
2. Download the English words file:
```bash
make download-words
```

## Usage

Run with default settings:
```bash
go run cmd/main.go
```

Run with command-line flags:
```bash
go run cmd/main.go -num 1000 -min 5 -max 500 -avg 100 -output output.txt
```

Run in interactive mode:
```bash
go run cmd/main.go -interactive
```

### Command-line Flags

- `-num`: Number of documents to generate (default: 1,000,000)
- `-min`: Minimum words per document (default: 10)
- `-max`: Maximum words per document (default: 1,000)
- `-avg`: Target average words per document (default: 200)
- `-output`: Output file path (default: "documents.txt")
- `-words`: English words file path (default: "words.txt")
- `-interactive`: Run in interactive mode

## Project Structure

```
.
├── cmd/
│   └── main.go           # Main entry point
├── pkg/
│   ├── config/           # Configuration handling
│   │   └── config.go
│   ├── generator/        # Document generation logic
│   │   └── generator.go
│   └── utils/           # Utility functions
│       └── distribution.go
├── go.mod               # Go module file
├── README.md            # This file
└── words.txt            # Dictionary file (download separately)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
