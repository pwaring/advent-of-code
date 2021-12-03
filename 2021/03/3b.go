package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	oxygenGeneratorDiagnostics := []string{}

	// Read all diagnostic lines into slice
	for scanner.Scan() {
		oxygenGeneratorDiagnostics = append(oxygenGeneratorDiagnostics, scanner.Text())
	}

	// CO2 scrubber is the same as oxygen generator to begin with
	// We have to use copy() as assigning a slice to another variable creates a
	// reference to the slice rather than a copy
	// We also have to create space for the new slice using make() because copy()
	// will only copy the minimum of len(src) and len(dst)
	co2ScrubberDiagnostics := make([]string, len(oxygenGeneratorDiagnostics))
	copy(co2ScrubberDiagnostics, oxygenGeneratorDiagnostics)

	// ASSUMPTION: All diagnostics are of the same length
	diagnosticLength := len(oxygenGeneratorDiagnostics[0])

	for column := 0; column < diagnosticLength; column++ {
		oneCount := 0
		zeroCount := 0

		// Only check the diagnostics if there is more than one element remaining, otherwise
		// we might remove the final element
		if len(oxygenGeneratorDiagnostics) > 1 {
			for d := range oxygenGeneratorDiagnostics {
				// ASSUMPTION: All diagnostics are ASCII, one character = one byte
				if string(oxygenGeneratorDiagnostics[d][column]) == "1" {
					oneCount++
				} else {
					zeroCount++
				}
			}

			oxygenGeneratorKeep := ""

			if oneCount >= zeroCount {
				oxygenGeneratorKeep = "1"
			} else {
				oxygenGeneratorKeep = "0"
			}

			// Instead of trying to remove elements from existing slices, take a temporary
			// copy and then append the elements we want to keep to a new empty slice
			tmpDiagnostics := make([]string, len(oxygenGeneratorDiagnostics))
			copy(tmpDiagnostics, oxygenGeneratorDiagnostics)
			oxygenGeneratorDiagnostics = []string{}
			for t := range tmpDiagnostics {
				if string(tmpDiagnostics[t][column]) == oxygenGeneratorKeep {
					oxygenGeneratorDiagnostics = append(oxygenGeneratorDiagnostics, tmpDiagnostics[t])
				}
			}
		}

		// Now that we have finished the oxygen generators, do the same for the CO2 scrubbers
		oneCount = 0
		zeroCount = 0

		// Only check the diagnostics if there is more than one element remaining, otherwise
		// we might remove the final element
		if len(co2ScrubberDiagnostics) > 1 {
			for d := range co2ScrubberDiagnostics {
				// ASSUMPTION: All diagnostics are ASCII, one character = one byte
				if string(co2ScrubberDiagnostics[d][column]) == "1" {
					oneCount++
				} else {
					zeroCount++
				}
			}

			co2ScrubberKeep := ""
			if oneCount >= zeroCount {
				co2ScrubberKeep = "0"
			} else {
				co2ScrubberKeep = "1"
			}

			tmpDiagnostics := make([]string, len(co2ScrubberDiagnostics))
			copy(tmpDiagnostics, co2ScrubberDiagnostics)
			co2ScrubberDiagnostics = []string{}
			for t := range tmpDiagnostics {
				if string(tmpDiagnostics[t][column]) == co2ScrubberKeep {
					co2ScrubberDiagnostics = append(co2ScrubberDiagnostics, tmpDiagnostics[t])
				}
			}
		}
	}

	// ASSUMPTION: By the time we get to this point, the two diagnostics slices contain
	// exactly one element each
	oxygenGeneratorRating, _ := strconv.ParseInt(oxygenGeneratorDiagnostics[0], 2, 32)
	co2ScrubberRating, _ := strconv.ParseInt(co2ScrubberDiagnostics[0], 2, 32)

	lifeSupportRating := oxygenGeneratorRating * co2ScrubberRating

	fmt.Println(lifeSupportRating)
}
