package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	type InventoryItem struct {
		calories int
	}

	type InventoryGroup struct {
		items []InventoryItem
	}

	inventoryGroups := []InventoryGroup{}
	currentGroupIndex := 0
	currentGroupInitialised := false
	highestCalories := []int{0, 0, 0}

	scanner := bufio.NewScanner(os.Stdin)

	// Read in the data and convert into structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			if !currentGroupInitialised {
				inventoryGroups = append(inventoryGroups, InventoryGroup{})
				currentGroupInitialised = true
			}

			// ASSUMPTION: Every non-empty line contains a number which can be expressed as an integer
			var inventoryItem InventoryItem
			inventoryItem.calories, _ = strconv.Atoi(line)
			inventoryGroups[currentGroupIndex].items = append(inventoryGroups[currentGroupIndex].items, inventoryItem)
		} else {
			// Blank line separates groups
			currentGroupIndex++
			currentGroupInitialised = false
		}
	}

	// Now the data has been transformed into a structure, find the highest calorie group
	for groupIndex := range inventoryGroups {
		group := inventoryGroups[groupIndex]
		groupCalories := 0

		for itemIndex := range group.items {
			groupCalories += group.items[itemIndex].calories
		}

		// If this group has a higher calorie count than the lowest of the highest
		// calorie groups, replace the lowest with the current group
		if groupCalories > highestCalories[0] {
			highestCalories[0] = groupCalories

			// Resort the highest calories into ascending order, i.e. element 0
			// is the smallest
			sort.Slice(highestCalories, func(x, y int) bool {
				return highestCalories[x] < highestCalories[y]
			})
		}
	}

	// Finally, find the sum of the highest calories
	highestCaloriesSum := 0

	for hc := range highestCalories {
		highestCaloriesSum += highestCalories[hc]
	}

	fmt.Println(highestCaloriesSum)
}
