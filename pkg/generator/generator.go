package generator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"document-generator/pkg/config"
	"document-generator/pkg/utils"
)

// DocumentGenerator manages the generation of documents
type DocumentGenerator struct {
	Config     *config.Config
	Dictionary []string
	Rand       *rand.Rand
}

// New creates a new document generator with the given configuration
func New(config *config.Config) *DocumentGenerator {
	// Seed the random number generator
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	return &DocumentGenerator{
		Config: config,
		Rand:   r,
	}
}

// LoadDictionary loads the dictionary of words from the specified file
func (g *DocumentGenerator) LoadDictionary() error {
	fmt.Println("Loading English words from file...")
	words, err := loadEnglishWords(g.Config.WordsFile)
	if err != nil {
		return fmt.Errorf("error loading English words: %w", err)
	}

	g.Dictionary = words
	fmt.Printf("Loaded %d English words\n", len(g.Dictionary))
	return nil
}

// GenerateDocuments generates the specified number of documents and writes them to the output file
func (g *DocumentGenerator) GenerateDocuments() error {
	// Create output file
	file, err := os.Create(g.Config.OutputFile)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Use buffered writer for better performance
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Generate documents
	fmt.Printf("Generating %d documents...\n", g.Config.NumDocuments)

	// Track progress
	progressInterval := g.Config.NumDocuments / 10
	if progressInterval == 0 {
		progressInterval = 1 // Ensure we show progress for small document counts
	}

	for i := 0; i < g.Config.NumDocuments; i++ {
		document := g.generateDocument()
		_, err := writer.WriteString(document)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}

		// Show progress
		if (i+1)%progressInterval == 0 || i+1 == g.Config.NumDocuments {
			fmt.Printf("Progress: %d/%d documents (%.1f%%)\n",
				i+1, g.Config.NumDocuments, float64(i+1)/float64(g.Config.NumDocuments)*100)
		}
	}

	// Ensure all data is written
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing data: %w", err)
	}

	return nil
}

// generateDocument creates a single document with a random number of words
func (g *DocumentGenerator) generateDocument() string {
	// Use a triangular distribution to get an average close to avgWords
	// while still having a range from minWords to maxWords
	numWords := utils.GetTriangularDistributedWordCount(
		g.Rand,
		g.Config.MinWordsPerDoc,
		g.Config.MaxWordsPerDoc,
		float64(g.Config.AvgWordsPerDoc),
	)

	document := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		document[i] = g.Dictionary[g.Rand.Intn(len(g.Dictionary))]
	}

	return fmt.Sprintf("%s\n", utils.JoinWords(document))
}

// loadEnglishWords loads real English words from the specified file
func loadEnglishWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open words file: %w", err)
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
		return nil, fmt.Errorf("error reading words file: %w", err)
	}

	if len(words) == 0 {
		return nil, fmt.Errorf("no words found in file")
	}

	return words, nil
}
