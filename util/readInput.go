package util

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ReadInput(path string) []string {
	var lines []string
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err != nil {
		log.Fatal(err)
	}

	return lines
}
