package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time           int
	recordDistance int
}

func getRaces(str string) []Race {
	races := []Race{}

	parts := strings.Split(str, "\n")
	times := strings.Fields(strings.TrimSpace(parts[0]))
	recordDistances := strings.Fields(strings.TrimSpace(parts[1]))

	// Assumption: Times and record distances have the same number of fields
	// We skip the first index as that will be the label
	for index := 1; index < len(times); index++ {
		race := Race{}
		race.time, _ = strconv.Atoi(times[index])
		race.recordDistance, _ = strconv.Atoi(recordDistances[index])
		races = append(races, race)
	}

	return races
}

func getRaceWinningOptionsCount(race Race) int {
	winningOptionsCount := 0

	for holdTime := 0; holdTime <= race.time; holdTime++ {
		travelTime := race.time - holdTime
		distance := travelTime * holdTime

		if distance > race.recordDistance {
			winningOptionsCount++
		}
	}

	return winningOptionsCount
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	races := getRaces(inputString)

	winningProduct := 1

	for r := range races {
		winningProduct = winningProduct * getRaceWinningOptionsCount(races[r])
	}

	fmt.Println(winningProduct)
}
