package main

import (
	"awesomeProject/util"
	"fmt"
	"regexp"
	"strconv"
	strings "strings"
)

type Passport map[string]string

func main() {
	linesOfInputFile := util.ReadInput("day4/input.txt")
	passports := getPassportsFromInput(linesOfInputFile)
	validPassports := 0
	for _, passport := range passports {
		if passportValid(passport) {
			validPassports++
		} else {
			print(passport)
		}
	}

	fmt.Printf("In total we found %d valid passwords", validPassports)
}

func passportHasAllFields(passport Passport) bool {
	_, isNormalPassport := passport["cid"]
	return len(passport) == 8 || (len(passport) == 7 && !isNormalPassport)
}

func passportValid(passport Passport) bool {
	return passportHasAllFields(passport) &&
		validByr(passport) &&
		validIyr(passport) &&
		validEyr(passport) &&
		validHgt(passport) &&
		validHcl(passport) &&
		validEcl(passport) &&
		validPid(passport)
}

func getPassportsFromInput(input []string) []Passport {
	passportStrings := strings.Split(strings.Join(input, " "), "  ")
	var passports []Passport
	for _, passportString := range passportStrings {
		passports = append(passports, getPassportFromString(passportString))
	}
	return passports
}

func getPassportFromString(passportString string) Passport {
	passport := make(map[string]string)

	for _, keyValue := range strings.Fields(passportString) {
		keyValueArray := strings.SplitN(keyValue, ":", 2)
		passport[keyValueArray[0]] = keyValueArray[1]
	}

	return passport
}

func validByr(passport Passport) bool {
	byr, _ := strconv.Atoi(passport["byr"])
	return 1920 <= byr && byr <= 2002
}

func validIyr(passport Passport) bool {
	iyr, _ := strconv.Atoi(passport["iyr"])
	return 2010 <= iyr && iyr <= 2020
}

func validEyr(passport Passport) bool {
	eyr, _ := strconv.Atoi(passport["eyr"])
	return 2020 <= eyr && eyr <= 2030
}

func validHgt(passport Passport) bool {
	var (
		height int
		unit   string
	)
	fmt.Fscanf(strings.NewReader(passport["hgt"]), "%d%s", &height, &unit)

	if unit == "cm" {
		return 150 <= height && height <= 193
	} else if unit == "in" {
		return 59 <= height && height <= 76
	} else {
		return false
	}
}

func validHcl(passport Passport) bool {
	hcl, _ := passport["hcl"]
	var isValidHcl = regexp.MustCompile(`^#[a-f0-9]{6}$`).MatchString
	return isValidHcl(hcl)
}

func validEcl(passport Passport) bool {
	ecl, _ := passport["ecl"]
	var isValidEcl = regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)$`).MatchString
	return isValidEcl(ecl)
}

func validPid(passport Passport) bool {
	pid, _ := passport["pid"]
	var isValidPid = regexp.MustCompile(`^\d{9}$`).MatchString
	return isValidPid(pid)
}
