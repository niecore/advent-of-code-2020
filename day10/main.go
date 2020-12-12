package main

import (
	"../util"
	"sort"
)

func main() {
	adapters := util.ReadInputAsInts("day10/input.txt")

	joltDifferenceJumps := calculateJoltageJumps(adapters)
	println(util.CountItem(joltDifferenceJumps, 1) * util.CountItem(joltDifferenceJumps, 3))

	adapterCombinations := calculateAdapterCombinations(adapters)
	println(adapterCombinations)
}

func calculateAdapterCombinations(adaptersInBag []int) int {
	adapters := append([]int{0}, adaptersInBag...)
	joltageJumps := calculateJoltageJumps(adapters)

	combinationsToConnectAdapter := make([]int, len(joltageJumps))

	for i := len(joltageJumps) - 1; i >= 0; i-- {
		combinations := 0
		for j := i; j > (i - 3); j-- {
			if (j > 0 && joltageJumps[j] == 1) || j < 0 {
				combinations++
			} else {
				break
			}
		}

		if combinations == 0 {
			combinations = 1
		}
		combinationsToConnectAdapter[i] = combinations
	}

	combinationsToConnectAdapterSum := make([]int, len(joltageJumps))
	combinationsToConnectAdapterSum[0] = 1

	for index, combination := range combinationsToConnectAdapter[1:] {

		if combination == 1 {
			combinationsToConnectAdapterSum[index+1] = combinationsToConnectAdapterSum[index]
		} else if combination == 2 {
			combinationsToConnectAdapterSum[index+1] = combinationsToConnectAdapterSum[index] + combinationsToConnectAdapterSum[index-1]
		} else if combination == 3 {
			combinationsToConnectAdapterSum[index+1] = combinationsToConnectAdapterSum[index] + combinationsToConnectAdapterSum[index-1] + combinationsToConnectAdapterSum[index-2]
		}
	}

	return combinationsToConnectAdapterSum[len(joltageJumps)-1]

}

func calculateJoltageJumps(adapters []int) []int {
	currentJoltage := 0

	_, targetJoltage := util.FindMinMax(adapters)
	targetJoltage += 3

	adapters = append(adapters, targetJoltage)
	sort.Ints(adapters)

	joltageJumps := make([]int, len(adapters))

	for index, outputJoltage := range adapters {
		difference := outputJoltage - currentJoltage
		currentJoltage = outputJoltage
		joltageJumps[index] = difference
	}

	return joltageJumps
}
