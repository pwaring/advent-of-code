package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type Crate struct {
		id string
	}

	type Stack struct {
		crates []Crate
	}

	type Instruction struct {
		count  int
		source int
		target int
	}

	instructions := []Instruction{}
	stacks := []Stack{}
	lines := []string{}

	// Add an empty zero-indexed stack so we start from 1 when adding
	// real stacks
	emptyStack := Stack{}
	stacks = append(stacks, emptyStack)

	scanner := bufio.NewScanner(os.Stdin)

	// Read the file into a slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Reverse the slice of lines so we can iterate over it backwards
	for x, y := 0, len(lines)-1; x < y; x, y = x+1, y-1 {
		lines[x], lines[y] = lines[y], lines[x]
	}

	for lineIndex := range lines {
		line := lines[lineIndex]

		if strings.Index(line, "move") == 0 {
			// Instruction
			parts := strings.Fields(line)
			instruction := Instruction{}
			instruction.count, _ = strconv.Atoi(parts[1])
			instruction.source, _ = strconv.Atoi(parts[3])
			instruction.target, _ = strconv.Atoi(parts[5])
			instructions = append(instructions, instruction)
		} else if strings.Contains(line, "[") {
			// Crates
			// Each crate definition is 4 characters
			// ASSUMPTION: By the time we process the crate definitions,
			// we have built all the stacks
			for start, end, stackNumber := 0, 3, 1; start < len(line); start, end, stackNumber = start+4, end+4, stackNumber+1 {
				crateDefinition := line[start:end]

				if strings.Contains(crateDefinition, "[") {
					crate := Crate{}
					crate.id = crateDefinition[1:2]
					stacks[stackNumber].crates = append(stacks[stackNumber].crates, crate)
				}
			}
		} else if strings.Index(strings.TrimSpace(line), "1") == 0 {
			// Stacks
			// ASSUMPTION: Stack numbers will be sequential and start from 1
			stackNumbers := strings.Fields(line)

			for s := 0; s < len(stackNumbers); s++ {
				stack := Stack{}
				stacks = append(stacks, stack)
			}
		}
	}

	// We now have all the stacks, so process the instructions
	// Reverse the slice of instructions as they will be in the wrong order
	for x, y := 0, len(instructions)-1; x < y; x, y = x+1, y-1 {
		instructions[x], instructions[y] = instructions[y], instructions[x]
	}

	for i := range instructions {
		moveCrates := []Crate{}

		for c := 1; c <= instructions[i].count; c++ {
			// Pop top crate from source stack - this requires
			// two assignments in Go because there is no slice.pop
			topCrate := stacks[instructions[i].source].crates[len(stacks[instructions[i].source].crates)-1]
			stacks[instructions[i].source].crates = stacks[instructions[i].source].crates[:len(stacks[instructions[i].source].crates)-1]

			moveCrates = append(moveCrates, topCrate)
		}

		// Reverse the crates to be moved so we add them in the correct order
		for x, y := 0, len(moveCrates)-1; x < y; x, y = x+1, y-1 {
			moveCrates[x], moveCrates[y] = moveCrates[y], moveCrates[x]
		}

		// Add each crate one at a time
		for mc := range moveCrates {
			// Move the top crate to the target stack
			stacks[instructions[i].target].crates = append(stacks[instructions[i].target].crates, moveCrates[mc])
		}
	}

	// Print the top crate from each stack (if there is one)
	for s := range stacks {
		if len(stacks[s].crates) >= 1 {
			fmt.Print(stacks[s].crates[len(stacks[s].crates)-1].id)
		}
	}

	fmt.Println("")
}
