package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

const (
	Unknown = iota
	Operational
	Damaged
)

type SpringsRow struct {
	springs       []int
	damagedGroups []int
}

func powInt(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func populatePlaceholders(springs []int, values []int) []int {
	springsCopy := make([]int, len(springs))
	copy(springsCopy, springs)

	for sc, v := 0, 0; sc < len(springsCopy) && v < len(values); sc++ {
		if springsCopy[sc] == Unknown {
			springsCopy[sc] = values[v]
			v++
		}
	}

	return springsCopy
}

func calculateDamagedGroups(springs []int) []int {
	damagedGroups := []int{}

	for s, damageCount := 0, 0; s < len(springs); s++ {
		if springs[s] == Damaged {
			damageCount++

			// Edge case: last spring
			if s == len(springs)-1 {
				damagedGroups = append(damagedGroups, damageCount)
			}
		} else {
			if damageCount > 0 {
				damagedGroups = append(damagedGroups, damageCount)
				damageCount = 0
			}
		}
	}

	return damagedGroups
}

func getCombinations(choices []int, placeholderCount int) [][]int {
	choiceCount := len(choices)
	columnCount := placeholderCount
	rowCount := powInt(choiceCount, placeholderCount)
	combinations := [][]int{}

	// Create the grid of rows and columns
	for row := 0; row < rowCount; row++ {
		combination := make([]int, columnCount)
		combinations = append(combinations, combination)
	}

	for column := 0; column < columnCount; column++ {
		repeatCount := rowCount / powInt(choiceCount, column+1)
		choiceIndex := 0
		for row := 0; row < rowCount; {
			// Repeat choices for count
			// Repeat the repeats until we have filled all the rows
			for rc := 1; rc <= repeatCount && row < rowCount; rc++ {
				combinations[row][column] = choices[choiceIndex]
				row++
			}

			if choiceIndex == len(choices)-1 {
				choiceIndex = 0
			} else {
				choiceIndex++
			}
		}
	}

	return combinations
}

func parseSpringsRows(str string) []SpringsRow {
	springsRows := []SpringsRow{}

	str = strings.TrimSpace(str)
	lines := strings.Split(str, "\n")

	for index := range lines {
		currentLine := strings.TrimSpace(lines[index])

		if len(currentLine) > 0 {
			springsRows = append(springsRows, parseSpringsRow(currentLine))
		}
	}

	return springsRows
}

func parseSpringsRow(str string) SpringsRow {
	springsRow := SpringsRow{}

	str = strings.TrimSpace(str)
	parts := strings.Fields(str)

	// Get the springs one character at a time
	springChars := strings.Split(parts[0], "")
	for sc := range springChars {
		switch springChars[sc] {
		case "?":
			springsRow.springs = append(springsRow.springs, Unknown)
		case ".":
			springsRow.springs = append(springsRow.springs, Operational)
		case "#":
			springsRow.springs = append(springsRow.springs, Damaged)
		}
	}

	// Get the damaged groups, separated by commas
	damagedGroups := strings.Split(parts[1], ",")
	for dg := range damagedGroups {
		number, _ := strconv.Atoi(damagedGroups[dg])
		springsRow.damagedGroups = append(springsRow.damagedGroups, number)
	}

	return springsRow
}

func drawSpringsRow(springsRow SpringsRow) string {
	str := ""

	for s := range springsRow.springs {
		switch springsRow.springs[s] {
		case Unknown:
			str += "?"
		case Operational:
			str += "."
		case Damaged:
			str += "#"
		}
	}

	str += " "

	str += strings.Join(
		lo.Map(springsRow.damagedGroups, func(x int, index int) string {
			return strconv.Itoa(x)
		}),
		",",
	)

	return str
}

func countPlaceholders(springs []int) int {
	placeholderCount := 0

	for s := range springs {
		if springs[s] == Unknown {
			placeholderCount++
		}
	}

	return placeholderCount
}

func countMatchingArrangements(springsRow SpringsRow, choices []int) int {
	matchingArrangements := 0

	combinations := getCombinations(choices, countPlaceholders(springsRow.springs))

	for c := range combinations {
		completedSprings := populatePlaceholders(springsRow.springs, combinations[c])

		if slices.Equal(calculateDamagedGroups(completedSprings), springsRow.damagedGroups) {
			matchingArrangements++
		}
	}

	return matchingArrangements
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	springsRows := parseSpringsRows(inputString)
	choices := []int{Operational, Damaged}

	matchingArrangements := 0

	for sr := range springsRows {
		matchingArrangements += countMatchingArrangements(springsRows[sr], choices)
	}

	fmt.Println(matchingArrangements)
}
