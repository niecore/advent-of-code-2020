package util

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func ReadInputAsInts(path string) []int {
	var retval []int
	for _, input := range ReadInput(path) {
		param, _ := strconv.ParseInt(input, 0, 32)
		retval = append(retval, int(param))
	}
	return retval
}
