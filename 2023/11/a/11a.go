package main

import (
	"fmt"
	"os"
	"strings"
)

type GalaxyLocation struct {
	x int
	y int
}

type GalaxyLocationPair struct {
	a        GalaxyLocation
	b        GalaxyLocation
	distance int
}

const (
	Space = iota
	Galaxy
)

func getTaxiCabDistance(x1 int, y1 int, x2 int, y2 int) int {
	// Go doesn't have an abs function for the int type so we have
	// to check which is the largest before doing the subtraction
	xDiff := 0

	if x1 >= x2 {
		xDiff = x1 - x2
	} else {
		xDiff = x2 - x1
	}

	yDiff := 0

	if y1 >= y2 {
		yDiff = y1 - y2
	} else {
		yDiff = y2 - y1
	}

	return xDiff + yDiff
}

func pairGalaxyLocations(locations []GalaxyLocation) []GalaxyLocationPair {
	pairs := []GalaxyLocationPair{}

	// Pair order doesn't matter, e.g. if we have paired a + b
	// we don't also pair b + a
	// We don't check the last location because it will already
	// have been paired with every other location
	for a := 0; a < len(locations)-1; a++ {
		for b := a + 1; b < len(locations); b++ {
			pairs = append(pairs, GalaxyLocationPair{
				a: locations[a],
				b: locations[b],
				distance: getTaxiCabDistance(
					locations[a].x,
					locations[a].y,
					locations[b].x,
					locations[b].y,
				),
			})
		}
	}

	return pairs
}

func findGalaxyLocations(universe [][]int) []GalaxyLocation {
	galaxyLocations := []GalaxyLocation{}

	for row := range universe {
		for column := range universe[row] {
			if universe[row][column] == Galaxy {
				galaxyLocations = append(galaxyLocations, GalaxyLocation{
					x: column,
					y: row,
				})
			}
		}
	}

	return galaxyLocations
}

func drawUniverse(universe [][]int) string {
	str := ""

	for u := range universe {
		str += drawUniverseLine(universe[u]) + "\n"
	}

	return str
}

func drawUniverseLine(line []int) string {
	str := ""

	for index := range line {
		if line[index] == Galaxy {
			str += "#"
		} else if line[index] == Space {
			str += "."
		}
	}

	return str
}

func parseUniverse(str string) [][]int {
	universe := [][]int{}

	lines := strings.Split(str, "\n")

	for index := range lines {
		currentLine := strings.TrimSpace(lines[index])

		if len(currentLine) > 0 {
			universe = append(universe, parseUniverseLine(currentLine))
		}
	}

	return universe
}

func parseUniverseLine(line string) []int {
	parsedLine := []int{}

	tiles := strings.Split(strings.TrimSpace(line), "")

	for t := range tiles {
		if tiles[t] == "." {
			parsedLine = append(parsedLine, Space)
		} else if tiles[t] == "#" {
			parsedLine = append(parsedLine, Galaxy)
		}
	}

	return parsedLine
}

func expandUniverseRows(universe [][]int) [][]int {
	universeCopy := make([][]int, len(universe))
	copy(universeCopy, universe)
	expandedUniverse := [][]int{}

	for uRow := range universeCopy {
		expandedUniverse = append(expandedUniverse, universeCopy[uRow])

		// If row contains no galaxies, add it again
		if noGalaxies(universe[uRow]) {
			expandedUniverse = append(expandedUniverse, universeCopy[uRow])
		}
	}

	return expandedUniverse
}

func swapRowsColumns(universe [][]int) [][]int {
	universeCopy := make([][]int, len(universe))
	copy(universeCopy, universe)

	// First create swapped universe with empty cells
	swappedUniverse := [][]int{}

	// Create one row for each column
	rowCount := len(universeCopy)
	columnCount := len(universeCopy[0])
	for c := 1; c <= columnCount; c++ {
		swappedUniverse = append(swappedUniverse, make([]int, rowCount))
	}

	// Iterate over rows and columns and swap them
	for row := range universeCopy {
		for column := range universeCopy[row] {
			swappedUniverse[column][row] = universeCopy[row][column]
		}
	}

	return swappedUniverse
}

func expandUniverse(universe [][]int) [][]int {
	universeCopy := make([][]int, len(universe))
	copy(universeCopy, universe)

	// Expand rows first
	expandedUniverse := expandUniverseRows(universeCopy)

	// Swap rows and columns so we can expand columns as rows
	expandedUniverse = swapRowsColumns(expandedUniverse)
	expandedUniverse = expandUniverseRows(expandedUniverse)

	// Swap rows and columns to get back to original
	expandedUniverse = swapRowsColumns(expandedUniverse)

	return expandedUniverse
}

func noGalaxies(partialUniverse []int) bool {
	galaxyCount := 0

	for pu := range partialUniverse {
		if partialUniverse[pu] == Galaxy {
			galaxyCount++
		}
	}

	return galaxyCount == 0
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	universe := parseUniverse(inputString)
	expandedUniverse := expandUniverse(universe)

	galaxyLocations := findGalaxyLocations(expandedUniverse)
	galaxyPairs := pairGalaxyLocations(galaxyLocations)

	shortestPathSum := 0

	for gp := range galaxyPairs {
		shortestPathSum += galaxyPairs[gp].distance
	}

	fmt.Println(shortestPathSum)
}
