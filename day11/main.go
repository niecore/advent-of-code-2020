package main

import (
	"../util"
)

type SeatLayout struct {
	seats   []Seat
	rows    int
	columns int
}

type Seat struct {
	isSeat   bool
	occupied bool
}

func main() {
	oldSeatLayout := parseSeatLayout("day11/input.txt")

	for {
		newSeatLayout, changed := predictNewSeatLayout(oldSeatLayout, true)
		if !changed {
			break
		}

		printSeatLayout(newSeatLayout)
		oldSeatLayout = newSeatLayout
	}

	println(countOccupiedSeats(oldSeatLayout.seats))
}

func predictNewSeatLayout(oldSeatLayout SeatLayout, new bool) (SeatLayout, bool) {
	changed := false
	var newSeats []Seat
	for row := 0; row < oldSeatLayout.rows; row++ {
		for column := 0; column < oldSeatLayout.columns; column++ {
			seat, _ := getSeat(oldSeatLayout, row, column)
			newSeat := Seat{seat.isSeat, seat.occupied}

			var occupied int
			if new {
				occupied = occupiedVisibleSeats(oldSeatLayout, row, column)
			} else {
				occupied = occupiedAdjacentSeats(oldSeatLayout, row, column)
			}

			var occupiedLimit int
			if new {
				occupiedLimit = 5
			} else {
				occupiedLimit = 4
			}

			if seat.isSeat && seat.occupied && occupied >= occupiedLimit {
				newSeat.occupied = false
				changed = true
			} else if seat.isSeat && !seat.occupied && occupied == 0 {
				newSeat.occupied = true
				changed = true
			}

			newSeats = append(newSeats, newSeat)
		}
	}
	oldSeatLayout.seats = newSeats
	return oldSeatLayout, changed
}

func getSeat(layout SeatLayout, row int, column int) (Seat, bool) {
	if 0 <= row && row < layout.rows {
		if 0 <= column && column < layout.columns {
			return layout.seats[(row*layout.columns)+column], true
		}
	}
	return Seat{}, false
}

func getAdjacentSeat(layout SeatLayout, row int, column int, distance int) []Seat {
	var seats []Seat

	adjacent := [][]int{
		{row - distance, column - distance}, {row - distance, column}, {row - distance, column + distance},
		{row, column - distance}, {row, column + distance},
		{row + distance, column - distance}, {row + distance, column}, {row + distance, column + distance},
	}

	for _, acjecentCords := range adjacent {
		if seat, ok := getSeat(layout, acjecentCords[0], acjecentCords[1]); ok {
			seats = append(seats, seat)
		}
	}
	return seats
}

func occupiedAdjacentSeats(layout SeatLayout, row int, column int) int {
	occupiedCount := 0
	for _, seat := range getAdjacentSeat(layout, row, column, 1) {
		if seat.occupied {
			occupiedCount++
		}
	}
	return occupiedCount
}

func isOccupiedSeatInDirection(layout SeatLayout, row int, column int, x int, y int) bool {
	// up & down
	i := x
	j := y
	for {
		if seat, ok := getSeat(layout, row+i, column+j); ok {
			if seat.isSeat && !seat.occupied {
				return false
			} else if seat.isSeat && seat.occupied {
				return true
			}
		} else {
			return false
		}
		i += x
		j += y
	}
	return false
}

func occupiedVisibleSeats(layout SeatLayout, row int, column int) int {
	occupiedCount := 0

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, directions := range directions {
		if isOccupiedSeatInDirection(layout, row, column, directions[0], directions[1]) {
			occupiedCount++
		}
	}

	return occupiedCount
}

func countOccupiedSeats(seats []Seat) int {
	occupiedCount := 0
	for _, seat := range seats {
		if seat.occupied {
			occupiedCount++
		}
	}
	return occupiedCount
}

func printSeatLayout(layout SeatLayout) {
	for index, seat := range layout.seats {
		if index%layout.columns == 0 {
			println()
		}

		if !seat.isSeat {
			print(".")
		} else if seat.occupied {
			print("#")
		} else {
			print("L")
		}
	}
	println()
}

func parseSeatLayout(file string) SeatLayout {
	var seats []Seat
	rows := util.ReadInput(file)

	for _, row := range rows {
		for _, seat := range row {
			seats = append(seats, Seat{seat != '.', seat == '#'})
		}
	}
	return SeatLayout{seats: seats, rows: len(rows), columns: len(rows[0])}
}
