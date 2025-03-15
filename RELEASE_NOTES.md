# Document Generator v2025.03.14

## Overview

We're excited to announce the first release of Document Generator, a high-performance tool designed to generate large volumes of text documents using real English words. This tool is perfect for testing document processing systems, search engines, and other text-based applications that need realistic document datasets.

## Key Features

- **Configurable Document Generation**: Generate any number of documents with customizable length parameters
- **Real English Words**: Uses a comprehensive dictionary of English words for realistic content
- **Triangular Distribution**: Sophisticated algorithm ensures natural variation in document lengths while maintaining target averages
- **Progress Tracking**: Real-time progress updates during generation
- **Interactive Mode**: User-friendly interface for configuring generation parameters
- **High Performance**: Optimized for generating millions of documents efficiently

## Technical Details

- **Modular Architecture**: Clean separation of concerns with dedicated packages for configuration, generation, and utilities
- **Efficient I/O**: Uses buffered I/O for optimal performance when writing large files
- **Memory Efficient**: Reuses dictionary words to minimize memory usage
- **Configurable via CLI**: Comprehensive command-line interface for all settings
- **Well-Documented**: Includes detailed README and help commands

## Usage Examples

### Generate 1 Million Documents (Default)

```bash
make run
```

### Generate a Small Test Dataset

```bash
make run-small
```

### Interactive Configuration

```bash
make run-interactive
```

### Custom Configuration

```bash
./document-generator -num 50000 -min 20 -max 800 -avg 150 -output custom-docs.txt
```

## Installation

1. Clone the repository
2. Run `make download-words` to download the English dictionary
3. Run `make build` to compile the application

## Requirements

- Go 1.21 or higher
- Internet connection (for downloading the dictionary file)

## Future Plans

- Multi-threaded document generation for even faster performance
- Additional document formats and structures
- Statistical analysis of generated documents
- Custom word lists and dictionaries
- Document templates for more structured content

## Contributors

- Development Team

## License

MIT License 