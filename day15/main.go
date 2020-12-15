package main

func main() {
	startingNumbers := []int{9, 19, 1, 6, 0, 5, 4}
	memory := make(map[int]int)

	lastNumberSpoken := startingNumbers[0]
	for i := 1; i < len(startingNumbers); i++ {
		memory[lastNumberSpoken] = i - 1
		lastNumberSpoken = startingNumbers[i]
	}

	for i := len(startingNumbers); i < 30000000; i++ {
		lastTimeNumberWasSpoken, numberWasSpoken := memory[lastNumberSpoken]
		memory[lastNumberSpoken] = i - 1
		if numberWasSpoken {
			lastNumberSpoken = i - 1 - lastTimeNumberWasSpoken
		} else {
			lastNumberSpoken = 0
		}
	}
	println(lastNumberSpoken)
}
