package utilities

import (
	"log"
	"os"
	"strings"
)

// ReadInputFileToLines parses a file at filePath into its individual lines
func ReadInputFileToLines(filePath string) []string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(file), "\n")
}

// ReadInputFileToCharacters parses a file at filePath into individual characters
// this assumes the input file is a single line
func ReadInputFileToCharacters(filePath string) []string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(file), "")
}
