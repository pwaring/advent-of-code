package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	diagnostics := []string{}

	// Read all diagnostic lines into slice
	for scanner.Scan() {
		diagnostics = append(diagnostics, scanner.Text())
	}

	// Process diagnostics
	// ASSUMPTION: All diagnostics are of the same length
	diagnosticLength := len(diagnostics[0])
	gammaRateStr := ""
	epsilonRateStr := ""

	for column := 0; column < diagnosticLength; column++ {
		oneCount := 0
		zeroCount := 0

		for d := range diagnostics {
			// ASSUMPTION: All diagnostics are ASCII, one character = one byte
			if string(diagnostics[d][column]) == "1" {
				oneCount++
			} else {
				zeroCount++
			}
		}

		if oneCount > zeroCount {
			gammaRateStr += "1"
			epsilonRateStr += "0"
		} else {
			gammaRateStr += "0"
			epsilonRateStr += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateStr, 2, 32)
	epsilonRate, _ := strconv.ParseInt(epsilonRateStr, 2, 32)

	powerConsumption := gammaRate * epsilonRate

	fmt.Println(powerConsumption)
}
