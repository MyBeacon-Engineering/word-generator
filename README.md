# Document Generator

A Go program that generates a large text file (`documents.txt`) containing random documents.

## Features

- Generates 1 million documents in a single text file
- Each document is on its own line
- Documents contain between 10-1000 random words
- Average document length is 200 words
- Uses a dictionary of 10,000 random words for better performance
- Shows progress during generation

## Usage

To run the document generator:

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
- `dictionarySize`: Number of unique random words to generate (default: 10,000)
- `minWordLength`: Minimum length of each random word (default: 3)
- `maxWordLength`: Maximum length of each random word (default: 12)

## Performance

The program uses buffered I/O for better performance when writing the large output file. It also pre-generates a dictionary of random words to avoid generating new random words for each document.

The triangular distribution is used to achieve the target average document length while maintaining the specified minimum and maximum document lengths.
