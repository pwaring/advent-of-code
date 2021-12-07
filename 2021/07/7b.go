package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type Crab struct {
		position int
	}

	crabs := []Crab{}
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		crabStrs := strings.Split(scanner.Text(), ",")

		for c := range crabStrs {
			crabPosition, _ := strconv.Atoi(crabStrs[c])
			crabs = append(crabs, Crab{
				position: crabPosition,
			})
		}
	}

	// Find the lowest and highest positions, as these are the lower and
	// upper bound of our search space
	// Start with the first crab
	lowestCrabPosition := crabs[0].position
	highestCrabPosition := crabs[0].position

	for c := range crabs {
		if crabs[c].position < lowestCrabPosition {
			lowestCrabPosition = crabs[c].position
		} else if crabs[c].position > highestCrabPosition {
			highestCrabPosition = crabs[c].position
		}
	}

	// Find the lowest fuel position and the expenditure required
	lowestFuelPosition := highestCrabPosition + 1
	lowestFuelExpenditure := -1

	for candidatePosition := lowestCrabPosition; candidatePosition <= highestCrabPosition; candidatePosition++ {
		positionFuelExpenditure := 0

		for c := range crabs {
			positionDifference := 0

			// Could use math.Abs, but that requires float64 whereas we are working with int
			if candidatePosition >= crabs[c].position {
				positionDifference = candidatePosition - crabs[c].position
			} else {
				positionDifference = crabs[c].position - candidatePosition
			}

			if positionDifference > 0 {
				// Use a specific example of Faulhaber's formula
				positionFuelExpenditure += (positionDifference + (positionDifference * positionDifference)) / 2
			}
		}

		if positionFuelExpenditure < lowestFuelExpenditure || lowestFuelExpenditure < 0 {
			lowestFuelExpenditure = positionFuelExpenditure
			lowestFuelPosition = candidatePosition
		}
	}

	fmt.Println("Position ", lowestFuelPosition, " gives the lowest fuel expenditure of: ", lowestFuelExpenditure)
}
