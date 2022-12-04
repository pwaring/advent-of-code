package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const compartmentCount = 2

	type RucksackCompartment struct {
		contents []string
	}

	type Rucksack struct {
		compartments [compartmentCount]RucksackCompartment
	}

	// IMPROVEMENT: This could be generated dynamically
	priorities := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}

	rucksacks := []Rucksack{}

	scanner := bufio.NewScanner(os.Stdin)

	// Read in the data and convert into structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		items := strings.Split(line, "")
		compartmentSize := len(items) / compartmentCount

		rucksack := Rucksack{}

		// ASSUMPTION: Compartments are always the same size and contents are
		// evenly distributed between compartments
		for compartmentIndex := 0; compartmentIndex < compartmentCount; compartmentIndex++ {
			compartmentStart := compartmentIndex * compartmentSize
			compartmentEnd := compartmentStart + compartmentSize
			rucksack.compartments[compartmentIndex].contents = items[compartmentStart:compartmentEnd]
		}

		rucksacks = append(rucksacks, rucksack)
	}

	// Search all rucksacks to find common items in compartments
	prioritiesSum := 0

	for rucksackIndex := range rucksacks {
		rucksackCompartments := rucksacks[rucksackIndex].compartments
		firstCompartment := rucksackCompartments[0]

		// Keep a track of items seen as there may be duplicates, but we only
		// want to count an item once
		seenItems := make(map[string]bool)

		for firstCompartmentContentIndex := range firstCompartment.contents {
			// Check if this item exists in all other compartments
			firstCompartmentItem := firstCompartment.contents[firstCompartmentContentIndex]
			itemCount := 1

			_, seen := seenItems[firstCompartmentItem]

			if !seen {
				seenItems[firstCompartmentItem] = true

				for compartmentIndex := 1; compartmentIndex < compartmentCount; compartmentIndex++ {
					matchFound := false
					for compartmentContentIndex := range rucksackCompartments[compartmentIndex].contents {
						if !matchFound && rucksackCompartments[compartmentIndex].contents[compartmentContentIndex] == firstCompartmentItem {
							itemCount++
							matchFound = true
						}
					}
				}

				if itemCount == compartmentCount {
					prioritiesSum += priorities[firstCompartmentItem]
				}
			}
		}
	}

	fmt.Println(prioritiesSum)
}
