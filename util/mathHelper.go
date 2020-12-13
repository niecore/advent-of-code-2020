package util

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func FindSum(list []int, sum int) bool {
	// https://web.stanford.edu/class/cs9/sample_probs/TwoSum.pdf
	// can be improved
	for _, i := range list {
		for _, j := range list {
			if j+i == sum {
				return true
			}
		}
	}
	return false
}

func FindMinMax(numbers []int) (int, int) {
	max := 0
	min := 0xFFFFFFFFFFF

	for _, i := range numbers {
		if i > max {
			max = i
		}

		if i < min {
			min = i
		}
	}

	return min, max
}

func CountItem(list []int, sum int) int {
	count := 0
	for _, i := range list {
		if i == sum {
			count++
		}
	}
	return count
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
