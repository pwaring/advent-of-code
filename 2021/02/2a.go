package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	horizontalPosition := 0
	depth := 0
	commands := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	// Read all commands into slice
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	// Process commands
	for c := range commands {
		parts := strings.Fields(commands[c])
		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])

		if direction == "forward" {
			horizontalPosition += distance
		} else if direction == "down" {
			depth += distance
		} else if direction == "up" {
			depth -= distance
		}
	}

	result := horizontalPosition * depth

	fmt.Println(result)
}
