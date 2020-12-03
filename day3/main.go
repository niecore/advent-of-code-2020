package main

import (
	"../util"
	"fmt"
	"os"
)

type Position struct {
	X int
	Y int
}

type PositionDelta struct {
	X int
	Y int
}

func main() {
	linesOfInputFile := util.ReadInput("day3/input.txt")

	pos := Position{0, 0}
	delta := PositionDelta{1, 3}
	endPos :=	calculateFinalPosition(len(linesOfInputFile), pos, delta)
	treesHit := 0

	for {
		if positionIsTree(pos, linesOfInputFile) {
			treesHit++
		}

		pos = calculateNewPosition(pos, delta)

		if pos == endPos{break}
	}

	fmt.Printf("During your way you hit in total %d trees \n", treesHit)
	os.Exit(0)
}

func calculateNewPosition(pos Position, delta PositionDelta) Position {
	return Position{pos.X + delta.X, pos.Y + delta.Y}
}

func positionIsTree(pos Position, areaMap []string) bool {
		return areaMap[pos.X][pos.Y % len(areaMap[0])] == "#"[0]
}

func calculateFinalPosition(length int, pos Position, delta PositionDelta) Position {
	verticalSteps := length / (delta.X - pos.X)
	horizontalSteps := verticalSteps * delta.Y

	return Position{verticalSteps, horizontalSteps}
}
