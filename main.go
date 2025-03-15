package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Default values for command-line flags
const (
	defaultNumDocuments   = 1000000
	defaultMinWordsPerDoc = 10
	defaultMaxWordsPerDoc = 1000
	defaultAvgWordsPerDoc = 200
	defaultOutputFile     = "documents.txt"
	defaultWordsFile      = "words.txt"
)

// loadEnglishWords loads real English words from the specified file
func loadEnglishWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open words file: %v", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading words file: %v", err)
	}

	if len(words) == 0 {
		return nil, fmt.Errorf("no words found in file")
	}

	return words, nil
}

// generateDocument creates a document with a random number of words
// runInteractiveMode prompts the user for configuration values except wordsFile
func runInteractiveMode(numDocs, minWords, maxWords, avgWords int, outputFile, wordsFile string) (int, int, int, int, string, string) {
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
	input := getInput(fmt.Sprintf("Number of documents [%d]: ", numDocs), numDocs)
	fmt.Sscanf(input, "%d", &numDocs)
	
	// Minimum words per document
	input = getInput(fmt.Sprintf("Minimum words per document [%d]: ", minWords), minWords)
	fmt.Sscanf(input, "%d", &minWords)
	
	// Maximum words per document
	input = getInput(fmt.Sprintf("Maximum words per document [%d]: ", maxWords), maxWords)
	fmt.Sscanf(input, "%d", &maxWords)
	
	// Average words per document
	input = getInput(fmt.Sprintf("Target average words per document [%d]: ", avgWords), avgWords)
	fmt.Sscanf(input, "%d", &avgWords)
	
	// Output file
	outputFile = getInput(fmt.Sprintf("Output file [%s]: ", outputFile), outputFile)
	
	// Words file is no longer prompted for - using the default or command line value
	
	return numDocs, minWords, maxWords, avgWords, outputFile, wordsFile
}

func generateDocument(r *rand.Rand, dictionary []string, minWords, maxWords int, avgWords float64) string {
	// Use a triangular distribution to get an average close to avgWords
	// while still having a range from minWords to maxWords
	numWords := getTriangularDistributedWordCount(r, minWords, maxWords, avgWords)
	
	document := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		document[i] = dictionary[r.Intn(len(dictionary))]
	}
	
	return fmt.Sprintf("%s\n", joinWords(document))
}

// getTriangularDistributedWordCount returns a word count that follows a triangular distribution
// with the specified minimum, maximum, and mode (peak) values
func getTriangularDistributedWordCount(r *rand.Rand, min, max int, avg float64) int {
	// Calculate mode (peak) of triangular distribution to achieve desired average
	// For a triangular distribution, average = (min + max + mode) / 3
	// So mode = 3*avg - min - max
	mode := 3*avg - float64(min) - float64(max)
	
	// Ensure mode is within bounds
	if mode < float64(min) {
		mode = float64(min)
	} else if mode > float64(max) {
		mode = float64(max)
	}
	
	// Generate random number from triangular distribution
	u := r.Float64()
	f := (mode - float64(min)) / (float64(max) - float64(min))
	
	var result float64
	if u < f {
		result = float64(min) + math.Sqrt(u*(float64(max)-float64(min))*(mode-float64(min)))
	} else {
		result = float64(max) - math.Sqrt((1-u)*(float64(max)-float64(min))*(float64(max)-mode))
	}
	
	return int(result)
}

// joinWords joins words with spaces
func joinWords(words []string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}

func main() {
	// Define command-line flags
	numDocumentsPtr := flag.Int("num", defaultNumDocuments, "Number of documents to generate")
	minWordsPtr := flag.Int("min", defaultMinWordsPerDoc, "Minimum words per document")
	maxWordsPtr := flag.Int("max", defaultMaxWordsPerDoc, "Maximum words per document")
	avgWordsPtr := flag.Int("avg", defaultAvgWordsPerDoc, "Target average words per document")
	outputFilePtr := flag.String("output", defaultOutputFile, "Output file path")
	wordsFilePtr := flag.String("words", defaultWordsFile, "English words file path")
	interactivePtr := flag.Bool("interactive", false, "Run in interactive mode")
	
	// Parse command-line flags
	flag.Parse()
	
	// Configuration values
	numDocuments := *numDocumentsPtr
	minWordsPerDoc := *minWordsPtr
	maxWordsPerDoc := *maxWordsPtr
	avgWordsPerDoc := *avgWordsPtr
	outputFile := *outputFilePtr
	wordsFile := *wordsFilePtr
	
	// Interactive mode
	if *interactivePtr {
		numDocuments, minWordsPerDoc, maxWordsPerDoc, avgWordsPerDoc, outputFile, wordsFile = runInteractiveMode(
			numDocuments, minWordsPerDoc, maxWordsPerDoc, avgWordsPerDoc, outputFile, wordsFile)
		// Note: wordsFile is not prompted for in interactive mode, it uses the default or command line value
	}
	
	// Print configuration
	fmt.Println("Document Generator Configuration:")
	fmt.Printf("  Number of documents: %d\n", numDocuments)
	fmt.Printf("  Min words per document: %d\n", minWordsPerDoc)
	fmt.Printf("  Max words per document: %d\n", maxWordsPerDoc)
	fmt.Printf("  Target avg words per document: %d\n", avgWordsPerDoc)
	fmt.Printf("  Output file: %s\n", outputFile)
	fmt.Printf("  Words file: %s\n\n", wordsFile)
	
	startTime := time.Now()
	fmt.Println("Starting document generation...")
	
	// Seed the random number generator
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	
	// Load English words from file
	fmt.Println("Loading English words from file...")
	dictionary, err := loadEnglishWords(wordsFile)
	if err != nil {
		fmt.Printf("Error loading English words: %v\n", err)
		fmt.Println("Please run 'make download-words' to download the English words file.")
		return
	}
	
	fmt.Printf("Loaded %d English words\n", len(dictionary))
	
	// Create output file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	
	// Use buffered writer for better performance
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	
	// Generate documents
	fmt.Printf("Generating %d documents...\n", numDocuments)
	
	// Track progress
	progressInterval := numDocuments / 10
	
	for i := 0; i < numDocuments; i++ {
		document := generateDocument(r, dictionary, minWordsPerDoc, maxWordsPerDoc, float64(avgWordsPerDoc))
		_, err := writer.WriteString(document)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
		
		// Show progress
		if (i+1)%progressInterval == 0 || i+1 == numDocuments {
			fmt.Printf("Progress: %d/%d documents (%.1f%%)\n", 
				i+1, numDocuments, float64(i+1)/float64(numDocuments)*100)
		}
	}
	
	// Ensure all data is written
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing data: %v\n", err)
		return
	}
	
	elapsed := time.Since(startTime)
	fmt.Printf("Document generation completed in %s\n", elapsed)
	fmt.Printf("Output written to %s\n", outputFile)
}
