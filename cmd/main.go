package main

import (
	"fmt"
	"time"

	"document-generator/pkg/config"
	"document-generator/pkg/generator"
)

func main() {
	// Parse command-line flags
	cfg := config.ParseFlags()

	// Create document generator
	gen := generator.New(cfg)

	// Print configuration
	cfg.PrintConfig()

	startTime := time.Now()
	fmt.Println("Starting document generation...")

	// Load dictionary
	err := gen.LoadDictionary()
	if err != nil {
		fmt.Printf("%v\n", err)
		fmt.Println("Please run 'make download-words' to download the English words file.")
		return
	}

	// Generate documents
	err = gen.GenerateDocuments()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Document generation completed in %s\n", elapsed)
	fmt.Printf("Output written to %s\n", cfg.OutputFile)
}
