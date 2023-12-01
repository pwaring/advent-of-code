package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type Calibration struct {
		characters []string
		digits     []int
		value      int
	}

	calibrations := []Calibration{}

	scanner := bufio.NewScanner(os.Stdin)

	// Read in each line and convert into calibration data structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			characters := strings.Split(line, "")
			calibration := Calibration{
				characters: characters,
			}

			for c := range characters {
				if currentDigit, err := strconv.Atoi(characters[c]); err == nil {
					calibration.digits = append(calibration.digits, currentDigit)
				}
			}

			calibrations = append(calibrations, calibration)
		}
	}

	// Populate the value of each calibration
	for c := range calibrations {
		firstDigit := strconv.Itoa(calibrations[c].digits[0])
		lastDigit := strconv.Itoa(calibrations[c].digits[len(calibrations[c].digits)-1])

		calibrations[c].value, _ = strconv.Atoi(firstDigit + lastDigit)
	}

	// Sum the values
	sum := 0
	for c := range calibrations {
		sum += calibrations[c].value
	}

	fmt.Println(sum)
}
