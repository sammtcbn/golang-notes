package main

import (
	"fmt"
	"strings"
)

func countOccurrences(inputString, searchString string) int {
	count := strings.Count(inputString, searchString)
	return count
}

func main() {
	// Input string
	inputString := "This is an example string, hoping to find the substring we're looking for and count its occurrences. This example is a demonstration."

	// Substring to search for
	searchString := "is"

	// Count occurrences
	occurrences := countOccurrences(inputString, searchString)

	// Display the result
	fmt.Printf("The substring '%s' occurs %d times in the string.\n", searchString, occurrences)
}

