package main

import (
  "awesomeProject/util"
  "strconv"
  "strings"
)

type Instruction struct {
  name  string
  param int
}

func main() {
  instructions := parseInstructions("day8/input.txt")

  println(executeInstructions(instructions))

  for pc, instruction := range instructions {
    if instruction.name == "jmp" || instruction.name == "nop" {
      oldInstruction := instruction
      instructions[pc] = modifyInstruction(instruction)
      if acc, terminated := executeInstructions(instructions); terminated {
        println(acc)
        break
      }
      instructions[pc] = oldInstruction
    }
  }
}

func modifyInstruction(instruction Instruction) Instruction {
  if instruction.name == "jmp" {
    instruction.name = "nop"
  } else {
    instruction.name = "jmp"
  }
  return instruction
}

func parseInstructions(file string) []Instruction {
  var instructions []Instruction
  for _, instruction := range util.ReadInput(file) {
    instructionArray := strings.Split(instruction, " ")
    name := instructionArray[0]
    param, _ := strconv.ParseInt(instructionArray[1], 0, 32)
    instructions = append(instructions, Instruction{name, int(param)})
  }
  return instructions
}

func executeInstructions(instructions []Instruction) (int, bool) {
  acc := 0
  pc := 0
  end := len(instructions)
  endlessLoopDetection := make([]bool, end)

  for pc < end {
    if endlessLoopDetection[pc] {
      return acc, false
    } else {
      endlessLoopDetection[pc] = true
    }

    switch instructions[pc].name {
    case "jmp":
      pc += instructions[pc].param
    case "nop":
      pc += 1
    case "acc":
      acc += instructions[pc].param
      pc += 1
    }
  }
  return acc, true
}
