package main

import (
	"awesomeProject/util"
	"regexp"
	"strconv"
	"strings"
)

const FIELDS = 20

type FieldDescription struct {
	name  string
	area1 []int
	area2 []int
}

type Ticket [FIELDS]int

func main() {
	// parse input
	inputFile := util.ReadInput("day16/input.txt")
	fieldRestrictions := getFieldRestrictions(inputFile)
	myTicket := getMyTicket(inputFile)
	nearbyTickets := getNearbyTickets(inputFile)

	// part 1
	println(getScanningErrorRateOfNearbyTickets(nearbyTickets, fieldRestrictions))

	// part 2
	validNearbyTickets := getValidNearbyTickets(nearbyTickets, fieldRestrictions)
	orderedRestrictions := calculateFieldOrder(validNearbyTickets, fieldRestrictions)
	println(getMyDepartureDetails(myTicket, orderedRestrictions))
}

func getMyDepartureDetails(ticket Ticket, restrictions []FieldDescription) int {
	depatureDetails := 1
	for index, restriction := range restrictions {
		if strings.HasPrefix(restriction.name, "departure") {
			depatureDetails *= ticket[index]
		}
	}
	return depatureDetails
}

func calculateFieldOrder(tickets []Ticket, restrictions []FieldDescription) []FieldDescription {
	// create matrix
	matrix := make([][]bool, FIELDS)
	for i := 0; i < FIELDS; i++ {
		matrix[i] = make([]bool, FIELDS)
		for j := 0; j < FIELDS; j++ {
			matrix[i][j] = true
		}
	}

	// fill matrix with valid assignments
	for _, ticket := range tickets {
		for fieldIndex, field := range ticket {
			for restrictionIndex, restriction := range restrictions {
				if !fieldOk(field, restriction) {
					matrix[fieldIndex][restrictionIndex] = false
				}
			}
		}
	}

	// find which field can only match which description
	orderedRestrictions := make([]FieldDescription, FIELDS)
	usedHelper := make([]bool, FIELDS)
	for i := 0; i < FIELDS; i++ {
		for j := 0; j < FIELDS; j++ {
			if util.CountValid(matrix[j]) == i+1 {
				for k := 0; k < FIELDS; k++ {
					if matrix[j][k] && !usedHelper[k] {
						usedHelper[k] = true
						orderedRestrictions[j] = restrictions[k]
					}
				}
			}
		}
	}

	return orderedRestrictions
}

func getValidNearbyTickets(tickets []Ticket, restrictions []FieldDescription) []Ticket {
	var valid []Ticket
	for _, ticket := range tickets {
		if getScanningErrorOfTicket(ticket, restrictions) == 0 {
			valid = append(valid, ticket)
		}
	}
	return valid
}

func getScanningErrorRateOfNearbyTickets(tickets []Ticket, restrictions []FieldDescription) int {
	totalScanningErrorRate := 0

	for _, ticket := range tickets {
		totalScanningErrorRate += getScanningErrorOfTicket(ticket, restrictions)
	}

	return totalScanningErrorRate
}

func getMyTicket(file []string) Ticket {
	var ticket Ticket
	for i, line := range file {
		if strings.HasPrefix(line, "your ticket:") {
			return parseTicket(file[i+1])
		}
	}
	return ticket
}

func getScanningErrorOfTicket(ticket Ticket, fields []FieldDescription) int {
	scanningError := 0

	for _, fieldInTicket := range ticket {
		fieldValid := false
		for _, fieldDescription := range fields {
			if fieldOk(fieldInTicket, fieldDescription) {
				fieldValid = true
				break
			}
		}

		if !fieldValid {
			scanningError += fieldInTicket
		}
	}

	return scanningError
}

func fieldOk(value int, field FieldDescription) bool {
	return (field.area1[0] <= value && value <= field.area1[1]) || (field.area2[0] <= value && value <= field.area2[1])
}

func getNearbyTickets(file []string) []Ticket {
	var tickets []Ticket
	for i, line := range file {
		if strings.HasPrefix(line, "nearby tickets:") {
			for j := i + 1; j < len(file); j++ {
				tickets = append(tickets, parseTicket(file[j]))
			}
		}
	}
	return tickets
}

func parseTicket(ticketLine string) Ticket {
	var ticket Ticket
	for i, input := range strings.Split(ticketLine, ",") {
		param, _ := strconv.ParseInt(input, 0, 32)
		ticket[i] = int(param)
	}
	return ticket
}

func getFieldRestrictions(file []string) []FieldDescription {
	fieldDescriptions := make([]FieldDescription, FIELDS)
	input := strings.Join(file, "\n")
	restrictionRe := regexp.MustCompile(`(.*): (\d+)-(\d+) or (\d+)-(\d+)*`)

	for i, desc := range restrictionRe.FindAllStringSubmatch(input, -1) {
		name := desc[1]
		min1, _ := strconv.Atoi(desc[2])
		max1, _ := strconv.Atoi(desc[3])
		min2, _ := strconv.Atoi(desc[4])
		max2, _ := strconv.Atoi(desc[5])

		fieldDescriptions[i] = FieldDescription{name, []int{min1, max1}, []int{min2, max2}}
	}

	return fieldDescriptions
}
