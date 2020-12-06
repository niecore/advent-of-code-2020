package main

import (
	"../util"
	"fmt"
	strings "strings"
)

func main() {
	linesOfInputFile := util.ReadInput("day6/input.txt")
	surveyStrings := strings.Split(strings.Join(linesOfInputFile, " "), "  ")

	questionsAnsweredByAnyone := 0
	questionsAnsweredByEveryone := 0

	for _, groupString := range surveyStrings {

		groupSize := len(strings.Split(groupString, " "))
		answersOfGroupList := strings.ReplaceAll(groupString, " ", "")

		answersOfGroup := map[rune]int{}

		for _, question := range answersOfGroupList {
			answersOfGroup[question]++
		}

		for _, numbOfYesAnswers := range answersOfGroup {
			if numbOfYesAnswers == groupSize {
				questionsAnsweredByEveryone++
			}
		}
		questionsAnsweredByAnyone += len(answersOfGroup)
	}

	fmt.Printf("Total answered questions %d\n", questionsAnsweredByAnyone)
	fmt.Printf("Everbody answered questions %d\n", questionsAnsweredByEveryone)
}
