package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "path/filepath"
  "strconv"
)

func readInput() []int {
  var lines []int
  absPath, _ := filepath.Abs("input.txt")
  file, err := os.Open(absPath)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line, _ := strconv.Atoi(scanner.Text())
    lines = append(lines, line)
  }

  if err != nil {
    log.Fatal(err)
  }

  return lines
}

func numberInList(a int, list []int) bool {
  for _, b := range list {
    if b == a {
      return true
    }
  }
  return false
}

func main() {
  numbers := readInput()

  for _, a := range numbers {
    b := 2020 - a
    if numberInList(b, numbers) {
      fmt.Printf("Found the correct numbers in the expenses report: %d, %d with the result %d", a, b, a*b)
      os.Exit(0)
    }
  }

  os.Exit(1)
}
