package main

import (
	"awesomeProject/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	linesOfInputFile := util.ReadInput("day5/input.txt")

	ticketIds := getTicketIds(linesOfInputFile)
	getHighestTicketId(ticketIds)
	findFreePlace(ticketIds)
}

func findFreePlace(ids []int) {
	sort.Ints(ids)
	bookedPlace := ids[0]

	for _, place := range ids[1:] {
		if bookedPlace+1 == place {
			bookedPlace = place
			continue
		} else {
			println("Your free seat is: %d", place-1)
			break
		}
	}
}

func getHighestTicketId(ticketIds []int) {
	highestTicketId := 0

	for _, id := range ticketIds {
		if id > highestTicketId {
			highestTicketId = id
		}
	}

	fmt.Printf("The highest ticket id is %d\n", highestTicketId)
}

func getTicketIds(linesOfInputFile []string) []int {
	var tickets []int
	for _, ticket := range linesOfInputFile {
		var rowString string
		var columnString string

		fmt.Fscanf(strings.NewReader(ticket), "%7s%3s", &rowString, &columnString)

		row := getRowFromString(rowString)
		column := getColumnFromString(columnString)
		id := calculateSeatId(row, column)

		tickets = append(tickets, id)
	}

	return tickets
}

func calculateSeatId(row int, column int) int {
	return (row * 8) + column
}

func getRowFromString(row string) int {
	row = strings.Replace(row, "B", "1", -1)
	row = strings.Replace(row, "F", "0", -1)

	convertedRow, _ := strconv.ParseInt(row, 2, 8)
	convertedRowInt := int(convertedRow)
	return convertedRowInt
}

func getColumnFromString(column string) int {
	column = strings.Replace(column, "R", "1", -1)
	column = strings.Replace(column, "L", "0", -1)

	convertedColumn, _ := strconv.ParseInt(column, 2, 4)
	convertedColumnInt := int(convertedColumn)
	return convertedColumnInt
}
