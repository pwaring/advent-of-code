package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	increasedMeasurements := 0
	var input []int

	// Read file into slice, converting to integers
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		input = append(input, number)
	}

	// Calculate all the windows
	var windowSums []int
	windowSize := 3
	lastWindowIndex := len(input) - (windowSize - 1)

	for i := 0; i < lastWindowIndex; i++ {
		windowSums = append(windowSums, input[i]+input[i+1]+input[i+2])
	}

	// Count all the windows that are greater than the previous window
	for ws := range windowSums {
		if ws > 0 && windowSums[ws] > windowSums[ws-1] {
			increasedMeasurements++
		}
	}

	fmt.Println(increasedMeasurements)
}
