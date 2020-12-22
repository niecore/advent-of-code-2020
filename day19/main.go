package main

import (
	"../util"
	"strconv"
	"strings"
)

const SUBRULES = 2

type Rule struct {
	idx     int // debugging
	letter  string
	options [][]int
}

type Rules map[int]Rule

func main() {
	// parse input
	inputFile := util.ReadInput("day19/input.txt")

	rules := parseRules(inputFile)
	messages := parseMessages(inputFile)

	// remove for part 1
	rules[8] = Rule{idx: 8, options: [][]int{[]int{42}, []int{42, 8}}}
	rules[11] = Rule{idx: 11, options: [][]int{[]int{42, 31}, []int{42, 11, 31}}}
	// end remove

	count := 0
	for _, message := range messages {
		valid := isMessageValid(message, rules)
		if valid {
			count++
		}
	}
	println(count)
}

func isMessageValid(message string, rules Rules) bool {
	return checkRule(message, rules, []int{0})
}

func checkRule(message string, rules Rules, seq []int) bool {

	if len(message) == 0 && len(seq) == 0 {
		return true
	} else if len(message) == 0 || len(seq) == 0 {
		return false
	}

	rule0, _ := rules[seq[0]]
	newSeq := seq[1:]

	// this part applies only when we are looking at a base rules
	if rule0.letter != "" {
		if rule0.letter[0] == message[0] {
			return checkRule(message[1:], rules, newSeq)
		} else {
			return false
		}
	} else {
		for _, option := range rule0.options {

			if option == nil {
				continue
			}
			if checkRule(message, rules, append(option, newSeq...)) {
				return true
			}
		}
	}
	return false
}

func parseMessages(file []string) []string {
	for index, line := range file {
		if len(line) == 0 {
			return file[index+1:]
		}
	}
	return make([]string, 0)
}

func parseRules(file []string) Rules {
	rules := make(Rules)
	for _, line := range file {

		if len(line) == 0 {
			break
		}

		ruleArray := strings.Split(line, ": ")
		id, _ := strconv.Atoi(ruleArray[0])
		rule := ruleArray[1]

		ruleObj := parseRule(rule)
		ruleObj.idx = id
		rules[id] = ruleObj
	}
	return rules
}

func parseRule(ruleLine string) Rule {
	if strings.HasPrefix(ruleLine, "\"") {
		return Rule{letter: ruleLine[1 : len(ruleLine)-1]}
	} else {
		options := make([][]int, SUBRULES)
		for index, subRule := range strings.Split(ruleLine, "|") {
			subRules := strings.Split(strings.Trim(subRule, " "), " ")
			ruleArray := make([]int, len(subRules))
			for indexRule, stringRule := range subRules {
				value, _ := strconv.Atoi(stringRule)
				ruleArray[indexRule] = value
			}
			options[index] = ruleArray
		}

		return Rule{options: options}
	}
}
