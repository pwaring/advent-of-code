package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	overlapCount := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		assignments := strings.Split(line, ",")
		firstAssignment := strings.Split(assignments[0], "-")
		secondAssignment := strings.Split(assignments[1], "-")

		// Convert all assignments to integers, otherwise Go will compare them as strings
		// and this will not yield the result we expect
		a1, _ := strconv.Atoi(firstAssignment[0])
		a2, _ := strconv.Atoi(firstAssignment[1])
		b1, _ := strconv.Atoi(secondAssignment[0])
		b2, _ := strconv.Atoi(secondAssignment[1])

		if (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2) {
			overlapCount++
		}
	}

	fmt.Println(overlapCount)
}
