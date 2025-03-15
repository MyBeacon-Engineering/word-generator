package config

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Default configuration values
const (
	DefaultNumDocuments   = 1000000
	DefaultMinWordsPerDoc = 10
	DefaultMaxWordsPerDoc = 1000
	DefaultAvgWordsPerDoc = 200
	DefaultOutputFile     = "documents.txt"
	DefaultWordsFile      = "words.txt"
)

// Config holds all the configuration parameters for document generation
type Config struct {
	NumDocuments   int
	MinWordsPerDoc int
	MaxWordsPerDoc int
	AvgWordsPerDoc int
	OutputFile     string
	WordsFile      string
}

// NewDefaultConfig creates a new configuration with default values
func NewDefaultConfig() *Config {
	return &Config{
		NumDocuments:   DefaultNumDocuments,
		MinWordsPerDoc: DefaultMinWordsPerDoc,
		MaxWordsPerDoc: DefaultMaxWordsPerDoc,
		AvgWordsPerDoc: DefaultAvgWordsPerDoc,
		OutputFile:     DefaultOutputFile,
		WordsFile:      DefaultWordsFile,
	}
}

// ParseFlags parses command-line flags and returns a configuration
func ParseFlags() *Config {
	config := NewDefaultConfig()

	// Define command-line flags
	numDocumentsPtr := flag.Int("num", config.NumDocuments, "Number of documents to generate")
	minWordsPtr := flag.Int("min", config.MinWordsPerDoc, "Minimum words per document")
	maxWordsPtr := flag.Int("max", config.MaxWordsPerDoc, "Maximum words per document")
	avgWordsPtr := flag.Int("avg", config.AvgWordsPerDoc, "Target average words per document")
	outputFilePtr := flag.String("output", config.OutputFile, "Output file path")
	wordsFilePtr := flag.String("words", config.WordsFile, "English words file path")
	interactivePtr := flag.Bool("interactive", false, "Run in interactive mode")

	// Parse command-line flags
	flag.Parse()

	// Update configuration with flag values
	config.NumDocuments = *numDocumentsPtr
	config.MinWordsPerDoc = *minWordsPtr
	config.MaxWordsPerDoc = *maxWordsPtr
	config.AvgWordsPerDoc = *avgWordsPtr
	config.OutputFile = *outputFilePtr
	config.WordsFile = *wordsFilePtr

	// Interactive mode
	if *interactivePtr {
		promptForInput(config)
	}

	return config
}

// promptForInput prompts the user for configuration values in interactive mode
func promptForInput(config *Config) {
	reader := bufio.NewReader(os.Stdin)

	// Helper function to get user input with default value
	getInput := func(prompt string, defaultValue interface{}) string {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return fmt.Sprintf("%v", defaultValue)
		}
		// Trim whitespace and newline characters
		input = strings.TrimSpace(input)
		if input == "" {
			return fmt.Sprintf("%v", defaultValue)
		}
		return input
	}

	// Number of documents
	input := getInput(fmt.Sprintf("Number of documents [%d]: ", config.NumDocuments), config.NumDocuments)
	fmt.Sscanf(input, "%d", &config.NumDocuments)

	// Minimum words per document
	input = getInput(fmt.Sprintf("Minimum words per document [%d]: ", config.MinWordsPerDoc), config.MinWordsPerDoc)
	fmt.Sscanf(input, "%d", &config.MinWordsPerDoc)

	// Maximum words per document
	input = getInput(fmt.Sprintf("Maximum words per document [%d]: ", config.MaxWordsPerDoc), config.MaxWordsPerDoc)
	fmt.Sscanf(input, "%d", &config.MaxWordsPerDoc)

	// Average words per document
	input = getInput(fmt.Sprintf("Target average words per document [%d]: ", config.AvgWordsPerDoc), config.AvgWordsPerDoc)
	fmt.Sscanf(input, "%d", &config.AvgWordsPerDoc)

	// Output file
	config.OutputFile = getInput(fmt.Sprintf("Output file [%s]: ", config.OutputFile), config.OutputFile)
}

// PrintConfig prints the current configuration
func (c *Config) PrintConfig() {
	fmt.Println("Document Generator Configuration:")
	fmt.Printf("  Number of documents: %d\n", c.NumDocuments)
	fmt.Printf("  Min words per document: %d\n", c.MinWordsPerDoc)
	fmt.Printf("  Max words per document: %d\n", c.MaxWordsPerDoc)
	fmt.Printf("  Target avg words per document: %d\n", c.AvgWordsPerDoc)
	fmt.Printf("  Output file: %s\n", c.OutputFile)
	fmt.Printf("  Words file: %s\n\n", c.WordsFile)
}
