package main

import (
	"../util"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputLines := util.ReadInput("day13/input.txt")
	arrivalTime, _ := strconv.ParseInt(inputLines[0], 0, 32)
	busLines, offsets := parseBusLines(inputLines[1])

	println(calculateFastestArrivalOfBus(arrivalTime, busLines))
	println(calculateOffsetOfAlligningBus(busLines, offsets))
}

func calculateOffsetOfAlligningBus(busLines []int, offsets []int) int {
	valid := false
	t := 100000000000000
	for !valid {
		skip := 1
		valid = true
		for i := 0; i < len(busLines); i++ {
			offset := offsets[i]
			busId := busLines[i]

			if (t+offset)%busId != 0 {
				valid = false
				break
			}
			skip *= busId
		}

		if !valid {
			t += skip
		}
	}
	return t
}

func calculateFastestArrivalOfBus(time int64, busLines []int) int {
	waitTime := math.MaxInt32
	busLine := 0
	for _, _busLine := range busLines {
		_waitTime := _busLine - (int(time) % _busLine)
		if _waitTime < waitTime {
			waitTime = _waitTime
			busLine = _busLine
		}
	}

	return waitTime * busLine
}

func parseBusLines(s string) ([]int, []int) {
	var busLines []int
	var offset []int
	for index, busLine := range strings.Split(s, ",") {
		if busLine != "x" {
			busInterval, _ := strconv.ParseInt(busLine, 0, 32)
			busLines = append(busLines, int(busInterval))
			offset = append(offset, index)
		}
	}
	return busLines, offset
}
