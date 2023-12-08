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

func getRace(str string) Race {
	race := Race{}

	parts := strings.Split(str, "\n")
	times := strings.Fields(strings.TrimSpace(parts[0]))
	recordDistances := strings.Fields(strings.TrimSpace(parts[1]))

	time := ""
	recordDistance := ""

	// Assumption: Times and record distances have the same number of fields
	// We merge all the fields after the first into one numeric string
	for index := 1; index < len(times); index++ {
		time += times[index]
		recordDistance += recordDistances[index]
	}

	race.time, _ = strconv.Atoi(time)
	race.recordDistance, _ = strconv.Atoi(recordDistance)

	return race
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
	race := getRace(inputString)

	winningProduct := getRaceWinningOptionsCount(race)
	fmt.Println(winningProduct)
}
