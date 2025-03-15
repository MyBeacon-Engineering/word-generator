package utils

import (
	"math"
	"math/rand"
	"strings"
)

// GetTriangularDistributedWordCount returns a word count that follows a triangular distribution
// with the specified minimum, maximum, and mode (peak) values
func GetTriangularDistributedWordCount(r *rand.Rand, min, max int, avg float64) int {
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

// JoinWords joins words with spaces
func JoinWords(words []string) string {
	return strings.Join(words, " ")
}
