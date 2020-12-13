package main

import (
	"../util"
	"strconv"
)

type Instruction struct {
	direction string
	param     int
}

func main() {
	instructions := parseInstructions("day12/input.txt")

	println(executeInstructions(instructions))
}

func parseInstructions(file string) []Instruction {
	var instructions []Instruction
	for _, instruction := range util.ReadInput(file) {
		direction := instruction[0:1]
		param, _ := strconv.ParseInt(instruction[1:], 0, 32)
		instructions = append(instructions, Instruction{direction, int(param)})
	}
	return instructions
}

func executeInstructions(directions []Instruction) int {
	y := 0
	x := 0
	facing := "E"

	for _, direction := range directions {

		course, changeCourse := getNewDirection(facing, direction.direction, direction.param)

		if changeCourse {
			facing = course
		}

		if direction.direction == "R" || direction.direction == "L" {
			continue
		}

		switch course {
		case "E":
			x += direction.param
		case "W":
			x -= direction.param
		case "S":
			y -= direction.param
		case "N":
			y += direction.param
		}
	}
	return util.Abs(x) + util.Abs(y)
}

func getNewDirection(facing string, new string, degree int) (string, bool) {
	if new == "F" {
		return facing, false
	}

	if new == "E" || new == "W" || new == "S" || new == "N" {
		return new, false
	}

	turns := degree / 90
	tempFacing := facing
	for i := 0; i < turns; i++ {
		if new == "L" {
			tempFacing = turnLeft(tempFacing)
		} else {
			tempFacing = turnRight(tempFacing)
		}
	}
	return tempFacing, true
}

func turnLeft(facing string) string {
	switch facing {
	case "E":
		return "N"
	case "W":
		return "S"
	case "S":
		return "E"
	case "N":
		return "W"
	}
	return ""

}

func turnRight(facing string) string {
	switch facing {
	case "E":
		return "S"
	case "W":
		return "N"
	case "S":
		return "W"
	case "N":
		return "E"
	}
	return ""
}
