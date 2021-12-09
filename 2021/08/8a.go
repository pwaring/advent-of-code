package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SignalPattern struct {
	segments        map[string]bool
	segmentsOnCount int
}

type Display struct {
	input  [10]SignalPattern
	output [4]SignalPattern
}

func NewDisplay() *Display {
	var d Display

	for inputIndex := range d.input {
		d.input[inputIndex].segments = make(map[string]bool)
	}

	for outputIndex := range d.output {
		d.output[outputIndex].segments = make(map[string]bool)
	}

	return &d
}

func main() {
	uniqueSegments := map[int]bool{
		2: true,
		3: true,
		4: true,
		7: true,
	}

	displays := []Display{}

	// Read input and convert to data structures
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		inputPatterns := strings.Fields(strings.TrimSpace(parts[0]))
		outputPatterns := strings.Fields(strings.TrimSpace(parts[1]))

		tmpDisplay := NewDisplay()

		for ip := range inputPatterns {
			// ASSUMPTION: All patterns are ASCII, one character = one byte
			for c := 0; c < len(inputPatterns[ip]); c++ {
				character := string(inputPatterns[ip][c])
				tmpDisplay.input[ip].segments[character] = true
				tmpDisplay.input[ip].segmentsOnCount++
			}
		}

		for op := range outputPatterns {
			for c := 0; c < len(outputPatterns[op]); c++ {
				character := string(outputPatterns[op][c])
				tmpDisplay.output[op].segments[character] = true
				tmpDisplay.output[op].segmentsOnCount++
			}
		}

		displays = append(displays, *tmpDisplay)
	}

	uniqueSegmentCount := 0

	for displayIndex := range displays {
		display := displays[displayIndex]

		for outputIndex := range display.output {
			if uniqueSegments[display.output[outputIndex].segmentsOnCount] {
				uniqueSegmentCount++
			}
		}
	}

	fmt.Println("Outputs across all displays have the following number of unique segments: ", uniqueSegmentCount)
}
