package main

import (
	"../util"
	"strconv"
)

func main() {

	numbers := parseXmasCode("day9/input.txt")
	invalidNumber := getFirstInvalidNumber(numbers)

	n := 0
	m := 1
	for n <= m && m <= len(numbers) {
		if cur := util.Sum(numbers[n:m]); cur == invalidNumber {
			break
		} else if cur > invalidNumber {
			n++
		} else if cur < invalidNumber {
			m++
		}
	}
	min, max := util.FindMinMax(numbers[n:m])

	println(invalidNumber)
	println(min + max)
}

func getFirstInvalidNumber(numbers []int) int {
	pre := 25
	for i, sum := range numbers[pre:] {
		if !util.FindSum(numbers[i:i+pre], sum) {
			return sum
		}
	}
	return 0
}

func parseXmasCode(file string) []int {
	var xmasCode []int
	for _, input := range util.ReadInput(file) {
		param, _ := strconv.ParseInt(input, 0, 32)
		xmasCode = append(xmasCode, int(param))
	}
	return xmasCode
}
