package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Sequence struct {
	readings []int
}

type History struct {
	sequences []Sequence
}

func parseSequence(str string) Sequence {
	sequence := Sequence{}

	str = strings.TrimSpace(str)
	readings := strings.Fields(str)

	for r := range readings {
		number, _ := strconv.Atoi(readings[r])
		sequence.readings = append(sequence.readings, number)
	}

	return sequence
}

func isZero(sequence Sequence) bool {
	zeroCount := 0

	for r := range sequence.readings {
		if sequence.readings[r] == 0 {
			zeroCount++
		}
	}

	return zeroCount == len(sequence.readings)
}

func extrapolateSequence(sequence Sequence) []Sequence {
	sequences := []Sequence{}

	for currentSequence, zeroFound := sequence, false; !zeroFound; {
		sequences = append(sequences, currentSequence)

		if isZero(currentSequence) {
			zeroFound = true
		} else {
			nextSequence := Sequence{}

			// We work until the penultimate reading, because each reading
			// is compared to the next one
			for r := 0; r < len(currentSequence.readings)-1; r++ {
				nextSequence.readings = append(nextSequence.readings, currentSequence.readings[r+1]-currentSequence.readings[r])
			}

			currentSequence = nextSequence
		}
	}

	return sequences
}

func populateEstimates(sequences []Sequence) []Sequence {
	estimates := make([]Sequence, len(sequences))
	copy(estimates, sequences)

	// Estimates are calculated backwards from the last sequence
	lastEstimate := len(estimates) - 1
	for e := lastEstimate; e >= 0; e-- {
		if e == lastEstimate {
			// Last estimate is always zero
			estimates[e].readings = append(estimates[e].readings, 0)
		} else {
			// Other estimates are the last reading of the next estimate,
			// added to the last reading of the current estimate
			nextEstimate := &estimates[e+1]
			currentEstimate := &estimates[e]
			nextReading := nextEstimate.readings[len(nextEstimate.readings)-1] + currentEstimate.readings[len(currentEstimate.readings)-1]
			currentEstimate.readings = append(currentEstimate.readings, nextReading)
		}
	}

	return estimates
}

func main() {
	histories := []History{}
	scanner := bufio.NewScanner(os.Stdin)

	// Read in each line and convert to history structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			history := History{}
			history.sequences = populateEstimates(extrapolateSequence(parseSequence(line)))
			histories = append(histories, history)
		}
	}

	historySum := 0
	for h := range histories {
		firstReadings := histories[h].sequences[0].readings
		historySum += firstReadings[len(firstReadings)-1]
	}

	fmt.Println(historySum)
}
