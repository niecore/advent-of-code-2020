package main

import (
	"../util"
	"fmt"
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

	count := 0
	for _, message := range messages {
		valid := isMessageValid(message, rules)
		println(valid)
		if valid {
			count++
		}
	}
	println(count)
}

func isMessageValid(message string, rules Rules) bool {
	valid, parsed := checkRule(message, rules, rules[0], 0)
	return valid && parsed == len(message)
}

func checkRule(message string, rules Rules, rule Rule, pos int) (bool, int) {
	if rule.letter != "" {
		if pos <= len(message) && rule.letter[0] == message[pos] {
			return true, 1
		} else {
			return false, 0
		}
	}
	for _, option := range rule.options {
		parsed := 0
		valid := true
		for _, ruleIdx := range option {
			newRule := rules[ruleIdx]
			ok, parsedNew := checkRule(message, rules, newRule, pos+parsed)

			if !ok {
				valid = false
				break
			} else {
				parsed += parsedNew
			}
		}
		if valid {
			fmt.Printf("rule number %d parsed %d bytes to pos %d\n", rule.idx, parsed, pos+parsed)
			return true, parsed
		}
	}
	return false, 0
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
