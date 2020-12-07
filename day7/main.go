package main

import (
  "awesomeProject/util"
  "regexp"
  "strconv"
  "strings"
)

type BagParts map[string]int
type Bags map[string]BagParts

func main() {
  linesOfInputFile := util.ReadInput("day7/input.txt")

  bags := parseBags(linesOfInputFile)

  count := 0
  for bag, _ := range bags {
    if containsBagWithName("shiny gold", bag, bags) {
      count++
    }
  }
  println(count)                                   // part 1
  println(countTotalBagsInBag("shiny gold", bags)) // part 2
}

func containsBagWithName(find string, current string, bags Bags) bool {
  if _, bagIncluded := bags[current][find]; bagIncluded {
    return true
  } else {
    for bag, _ := range bags[current] {
      if containsBagWithName(find, bag, bags) {
        return true
      }
    }
    return false
  }
}

func countTotalBagsInBag(find string, bags Bags) int {
  if _, bagIncluded := bags[find]; !bagIncluded {
    return 0
  } else {
    bagCount := 0
    for bag, count := range bags[find] {
      bagCount += count                                  // count this bag
      bagCount += count * countTotalBagsInBag(bag, bags) // count sub bags of this abg
    }
    return bagCount
  }
}

func parseBags(file []string) Bags {
  bags := make(Bags)
  bagPartsRe := regexp.MustCompile(`(\d+)\s+([\s\w]+)\s+bags?[.|,]`)

  for _, bagDesc := range file {
    bagArray := strings.Split(bagDesc, " bags contain ")
    bagType := bagArray[0]
    bagParts := bagArray[1]

    bags[bagType] = make(BagParts)

    if bagParts == "no other bags" {
      continue
    } else {
      subBags := bagPartsRe.FindAllStringSubmatch(bagParts, -1)

      for _, subBag := range subBags {
        subBagCount, _ := strconv.Atoi(subBag[1])
        subBagType := subBag[2]

        bags[bagType][subBagType] = subBagCount
      }
    }
  }
  return bags
}
