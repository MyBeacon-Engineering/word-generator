package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	numDocuments     = 1000000
	minWordsPerDoc   = 10
	maxWordsPerDoc   = 1000
	avgWordsPerDoc   = 200
	outputFile       = "documents.txt"
	wordsFile        = "words.txt"
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
