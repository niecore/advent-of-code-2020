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

	combinationsToConnectAdapterSum := make([]int, len(joltageJumps))
	combinationsToConnectAdapterSum[0] = 1

	for i := 1; i < len(joltageJumps); i++ {
		// the adapter must be reached from the adapter in front
		combinationsToConnectAdapterSum[i] = combinationsToConnectAdapterSum[i-1]

		if joltageJumps[i] != 1 {
			// this adapter can only be reached with a 3 jump there is only a direct connection
			continue
		}

		for j := i - 1; j > (i - 3); j-- {
			if j > 0 && joltageJumps[j] == 1 {
				// the adapter on position J in front of us could be jumped
				// so we take the combinations of the adapter j-i also into consideration
				combinationsToConnectAdapterSum[i] += combinationsToConnectAdapterSum[j-1]
			} else {
				break
			}
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
