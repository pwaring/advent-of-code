package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type MoveInstruction struct {
	DeltaX int
	DeltaY int
}

func adjacent(a Point, b Point) bool {
	DeltaX := a.x - b.x
	DeltaY := a.y - b.y

	// Points are adjacent if their deltas in both directions are
	// at most one (overlapping counts as adjacent)
	return DeltaX >= -1 && DeltaX <= 1 && DeltaY >= -1 && DeltaY <= 1
}

func main() {
	tailVisits := make(map[Point]bool)

	head := Point{}
	tail := Point{}

	moveInstructions := []MoveInstruction{}

	// Read input and convert to data structures
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			fields := strings.Fields(line)
			direction := fields[0]
			steps, _ := strconv.Atoi(fields[1])

			// Break down multi-step instructions into series of
			// 1-step instructions for simplicity
			for s := 1; s <= steps; s++ {
				moveInstruction := MoveInstruction{}

				switch direction {
				case "L":
					moveInstruction.DeltaX = -1
				case "R":
					moveInstruction.DeltaX = 1
				case "U":
					moveInstruction.DeltaY = 1
				case "D":
					moveInstruction.DeltaY = -1
				}

				moveInstructions = append(moveInstructions, moveInstruction)
			}
		}
	}

	// Mark the starting point as visited by the tail
	tailVisits[tail] = true

	for mi := range moveInstructions {
		// Move the head
		head.x += moveInstructions[mi].DeltaX
		head.y += moveInstructions[mi].DeltaY

		if !adjacent(head, tail) {
			headTailDistanceX := tail.x - head.x
			headTailDistanceY := tail.y - head.y

			// Only move the tail along each axis if the distance is not
			// equal to zero
			if headTailDistanceX < 0 {
				tail.x += 1
			} else if headTailDistanceX > 0 {
				tail.x -= 1
			}

			if headTailDistanceY < 0 {
				tail.y += 1
			} else if headTailDistanceY > 0 {
				tail.y -= 1
			}

			tailVisits[tail] = true
		}
	}

	fmt.Println(len(tailVisits))
}
