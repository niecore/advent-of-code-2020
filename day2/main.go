package main

import (
	"../util"
	"fmt"
	"os"
	"strings"
)

func main() {
	passwordSpecifications := util.ReadInput("day2/input.txt")
	validPasswordsFound := 0
	for _, input := range passwordSpecifications {

		floor, ceil, letter, password, err := parsePasswordSpecification(input)

		if err != nil {
			fmt.Printf("Unable to parse password specification %s because %s\n", input, err)
			os.Exit(1)
		}

		if passwordIsValid(floor, ceil, letter, password) {
			validPasswordsFound++
		}
	}

	fmt.Printf("Found %d valid passwords\n", validPasswordsFound)
}

func parsePasswordSpecification(spec string) (int, int, string, string, error) {
	var (
		floor    int
		ceil     int
		letter   string
		password string
	)
	scanFormat := "%d-%d %1s: %s"
	_, err := fmt.Fscanf(strings.NewReader(spec), scanFormat, &floor, &ceil, &letter, &password)

	return floor, ceil, letter, password, err
}

func countLettersInPassword(letter string, password string) int {
	return strings.Count(password, letter)
}

func passwordIsValid(floor int, ceil int, letter string, password string) bool {
	letterOccurrence := countLettersInPassword(letter, password)
	return floor <= letterOccurrence && letterOccurrence <= ceil
}
