package main

import (
	"../util"
	"fmt"
	"os"
	"strings"
)

func main() {
	passwordSpecifications := util.ReadInput("day2/input.txt")
	validPasswordsFoundPart1 := 0
	validpasswordsFoundPart2 := 0
	for _, input := range passwordSpecifications {

		floor, ceil, letter, password, err := parsePasswordSpecification(input)

		if err != nil {
			fmt.Printf("Unable to parse password specification %s because %s\n", input, err)
			os.Exit(1)
		}

		if passwordIsValidPart1(floor, ceil, letter, password) {
			validPasswordsFoundPart1++
		}

		if passwordIsValidPart2(floor, ceil, letter, password) {
			validpasswordsFoundPart2++
		}
	}

	fmt.Printf("Found %d valid passwords for the first Policy\n", validPasswordsFoundPart1)
	fmt.Printf("Found %d valid passwords for the second Policy\n", validpasswordsFoundPart2)
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

func passwordIsValidPart1(floor int, ceil int, letter string, password string) bool {
	letterOccurrence := countLettersInPassword(letter, password)
	return floor <= letterOccurrence && letterOccurrence <= ceil
}

func passwordIsValidPart2(floor int, ceil int, letter string, password string) bool {
	return (password[floor-1] == letter[0]) != (password[ceil-1] == letter[0])
}
