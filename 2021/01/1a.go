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

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		input = append(input, number)
	}

	for i := range input {
		if i > 0 && input[i] > input[i-1] {
			increasedMeasurements++
		}
	}

	fmt.Println(increasedMeasurements)
}
