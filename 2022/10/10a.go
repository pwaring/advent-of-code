package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name     string
	argument int
}

func main() {
	instructions := []Instruction{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			fields := strings.Fields(line)
			instruction := Instruction{}
			instruction.name = fields[0]

			if instruction.name == "addx" {
				instruction.argument, _ = strconv.Atoi(fields[1])
			}

			instructions = append(instructions, instruction)
		}
	}

	cycles := []int{}
	cycles = append(cycles, 1)

	for i := range instructions {
		switch instructions[i].name {
		case "noop":
			cycles = append(cycles, cycles[len(cycles)-1])
		case "addx":
			// first cycle of addx is effectively a noop
			cycles = append(cycles, cycles[len(cycles)-1])
			cycles = append(cycles, cycles[len(cycles)-1]+instructions[i].argument)
		}
	}

	signalStrengthSum := 0

	for c := 20; c <= 220; c += 40 {
		// Value during a cycle is the value at the end of the previous cycle
		signalStrengthSum += c * cycles[c-1]
	}

	fmt.Println(signalStrengthSum)
}
