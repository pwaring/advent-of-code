package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToDigits(str string) []int {
	digitStrings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	digits := make([]int, 0)

	for c := 0; c < len(str); c++ {
		substr := str[c:]

		for ds := range digitStrings {
			if strings.HasPrefix(substr, ds) {
				digits = append(digits, digitStrings[ds])
			}
		}
	}

	return digits
}

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

			calibrations = append(calibrations, calibration)
		}
	}

	// Convert all characters into digits, including characters
	// that form a string which spells out a digit
	for calibrationIndex := range calibrations {
		currentString := ""

		// Keep reading in characters until we get to a numeric digit or the end of the string
		for characterIndex := 0; characterIndex < len(calibrations[calibrationIndex].characters); {
			currentChar := calibrations[calibrationIndex].characters[characterIndex]

			if numericDigit, err := strconv.Atoi(currentChar); err == nil {
				// Current character is a digit, so go back and process current string
				// to see if it contains any digits
				digits := stringToDigits(currentString)
				for d := range digits {
					calibrations[calibrationIndex].digits = append(calibrations[calibrationIndex].digits, digits[d])
				}

				// Add the numeric digit after any extracted from the current string
				calibrations[calibrationIndex].digits = append(calibrations[calibrationIndex].digits, numericDigit)
				characterIndex++
				currentString = ""
			} else if characterIndex == len(calibrations[calibrationIndex].characters)-1 {
				// End of string
				currentString += currentChar
				digits := stringToDigits(currentString)
				for d := range digits {
					calibrations[calibrationIndex].digits = append(calibrations[calibrationIndex].digits, digits[d])
				}
				characterIndex++
			} else {
				// Current character is not a digit, so add it to string
				currentString += currentChar
				characterIndex++
			}
		}
	}

	// Populate the value of each calibration
	for c := range calibrations {
		if len(calibrations[c].digits) > 0 {
			firstDigit := strconv.Itoa(calibrations[c].digits[0])
			lastDigit := strconv.Itoa(calibrations[c].digits[len(calibrations[c].digits)-1])

			calibrations[c].value, _ = strconv.Atoi(firstDigit + lastDigit)
		}
	}

	// Sum the values
	sum := 0
	for c := range calibrations {
		sum += calibrations[c].value
		fmt.Println(calibrations[c])
	}

	fmt.Println(sum)
}
