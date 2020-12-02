package main

import (
	"../util"
	"fmt"
	"os"
	"strconv"
)

func numberInList(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	linesOfInputFile := util.ReadInput("day1/input.txt")

	var numbers []int
	for _, x := range linesOfInputFile {
		number, _ := strconv.Atoi(x)
		numbers = append(numbers, number)
	}

	for _, a := range numbers {
		b := 2020 - a
		if numberInList(b, numbers) {
			fmt.Printf("Found the correct numbers in the expenses report: %d, %d with the result %d", a, b, a*b)
			os.Exit(0)
		}
	}

	os.Exit(1)
}
