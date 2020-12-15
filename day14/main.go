package main

import (
	"awesomeProject/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bitmask struct {
	zeros    uint64
	ones     uint64
	floating uint64
}

type Write struct {
	addr  uint64
	value uint64
}

type Memory map[uint64]uint64

func main() {

	var mask Bitmask
	memory := make(Memory)
	memory2 := make(Memory)

	for _, instruction := range util.ReadInput("day14/input.txt") {
		if isMask(instruction) {
			mask = parseBitmask(instruction)
		} else {
			write := parseWrite(instruction)
			executeMaskedValueWrite(memory, mask, write)
			executeMaskedAddressWrite(memory2, mask, write)
		}
	}

	println(getSumOfMemory(memory))
	println(getSumOfMemory(memory2))
}

func executeMaskedAddressWrite(memory Memory, mask Bitmask, write Write) {
	for _, address := range calculateAddresses(mask, write) {
		memory[address] = write.value
	}
}

func calculateAddresses(mask Bitmask, write Write) []uint64 {
	addresses := make([]uint64, 0)
	baseMask := (write.addr | mask.ones) & ^mask.floating
	addresses = append(addresses, baseMask)

	bitString := strconv.FormatInt(int64(mask.floating), 2)
	for index, floatingBit := range util.Reverse(bitString) {
		if floatingBit != '1' {
			continue
		}
		for _, address := range addresses {
			addresses = append(addresses, address|1<<index)
		}
	}
	return addresses
}

func executeMaskedValueWrite(memory Memory, mask Bitmask, write Write) {
	memory[write.addr] = getMaskedValue(mask, write.value)
}

func getMaskedValue(mask Bitmask, value uint64) uint64 {
	return (value | mask.ones) & ^mask.zeros
}

func isMask(instruction string) bool {
	return strings.HasPrefix(instruction, "mask")
}

func getSumOfMemory(memory Memory) uint64 {
	iv := uint64(0)

	for _, value := range memory {
		iv += value
	}

	return iv
}

func parseBitmask(bitmaskLine string) Bitmask {
	bitMaskString := strings.Split(bitmaskLine, " = ")[1]
	var zeros uint64
	var ones uint64
	var floating uint64
	for _, bit := range bitMaskString {
		zeros = zeros << 1
		ones = ones << 1
		floating = floating << 1
		switch bit {
		case '1':
			ones = ones | 1
		case '0':
			zeros = zeros | 1
		case 'X':
			floating = floating | 1
		}
	}

	return Bitmask{zeros: zeros, ones: ones, floating: floating}
}

func parseWrite(writeLine string) Write {
	var (
		addr  uint64
		value uint64
	)
	scanFormat := "mem[%d] = %d"
	_, err := fmt.Fscanf(strings.NewReader(writeLine), scanFormat, &addr, &value)

	if err != nil {
		os.Exit(1)
	}

	return Write{addr: addr, value: value}
}
