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

  deltas := []PositionDelta{
    PositionDelta{1, 1},
    PositionDelta{1, 3},
    PositionDelta{1, 5},
    PositionDelta{1, 7},
    PositionDelta{2, 1},
  }

  puzzleInput := 1

  for i, delta := range deltas {
    pos := Position{0, 0}
    endPos := calculateFinalPosition(len(linesOfInputFile)-1, pos, delta)
    treesHit := 0

    for {
      if positionIsTree(pos, linesOfInputFile) {
        treesHit++
      }

      if pos == endPos {
        break
      } else {
        pos = calculateNewPosition(pos, delta)
      }
    }
    fmt.Printf("During your way in iteration %d you hit in total %d trees \n", i, treesHit)
    puzzleInput = puzzleInput * treesHit
  }

  fmt.Printf("Your Puzzle Input is %d\n", puzzleInput)
  os.Exit(0)
}

func calculateNewPosition(pos Position, delta PositionDelta) Position {
  return Position{pos.X + delta.X, pos.Y + delta.Y}
}

func positionIsTree(pos Position, areaMap []string) bool {
  return areaMap[pos.X][pos.Y%len(areaMap[0])] == "#"[0]
}

func calculateFinalPosition(length int, pos Position, delta PositionDelta) Position {
  verticalSteps := length / (delta.X - pos.X)
  horizontalSteps := verticalSteps * delta.Y

  return Position{length, horizontalSteps}
}
