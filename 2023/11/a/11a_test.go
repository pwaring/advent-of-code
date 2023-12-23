package main

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

func TestParseUniverseLine(t *testing.T) {
	// Line 1 of test input
	assert.Equal(t, []int{
		Space,
		Space,
		Space,
		Galaxy,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
	}, parseUniverseLine("...#......"))

	// Line 10 of test input
	assert.Equal(t, []int{
		Galaxy,
		Space,
		Space,
		Space,
		Galaxy,
		Space,
		Space,
		Space,
		Space,
		Space,
	}, parseUniverseLine("#...#....."))

	// All spaces
	assert.Equal(t, []int{
		Space,
		Space,
		Space,
		Space,
		Space,
	}, parseUniverseLine("....."))

	// All galaxies
	assert.Equal(t, []int{
		Galaxy,
		Galaxy,
		Galaxy,
		Galaxy,
		Galaxy,
	}, parseUniverseLine("#####"))
}

func TestParseUniverse(t *testing.T) {
	universe := [][]int{}

	universe = append(universe, []int{
		Space,
		Space,
		Space,
		Galaxy,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Galaxy,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Galaxy,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Galaxy,
		Space,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Space,
		Galaxy,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Galaxy,
	})

	universe = append(universe, []int{
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Space,
		Galaxy,
		Space,
		Space,
	})

	universe = append(universe, []int{
		Galaxy,
		Space,
		Space,
		Space,
		Galaxy,
		Space,
		Space,
		Space,
		Space,
		Space,
	})

	assert.Equal(t, universe, parseUniverse(heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....
	`)))
}

func TestExpandUniverseRows(t *testing.T) {
	universe := parseUniverse(heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....
	`))

	expandedUniverse := parseUniverse(heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		..........
		......#...
		.#........
		.........#
		..........
		..........
		.......#..
		#...#.....
	`))

	assert.Equal(t, expandedUniverse, expandUniverseRows(universe))
}

func TestExpandUniverse(t *testing.T) {
	universe := parseUniverse(heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....
	`))

	expandedUniverse := parseUniverse(heredoc.Doc(`
		....#........
		.........#...
		#............
		.............
		.............
		........#....
		.#...........
		............#
		.............
		.............
		.........#...
		#....#.......
	`))

	assert.Equal(t, expandedUniverse, expandUniverse(universe))
}

func TestNoGalaxies(t *testing.T) {
	assert.True(t, noGalaxies([]int{
		Space,
		Space,
		Space,
		Space,
		Space,
	}))

	assert.False(t, noGalaxies([]int{
		Space,
		Space,
		Space,
		Space,
		Galaxy,
	}))

	assert.False(t, noGalaxies([]int{
		Galaxy,
		Galaxy,
		Galaxy,
		Galaxy,
		Galaxy,
	}))
}

func TestDrawUniverse(t *testing.T) {
	drawing := heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....
	`)

	assert.Equal(t, drawing, drawUniverse(parseUniverse(drawing)))
}

func TestSwapRowsColumns(t *testing.T) {
	universe := parseUniverse(heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....
	`))

	swappedUniverse := parseUniverse(heredoc.Doc(`
		..#......#
		.....#....
		..........
		#.........
		.........#
		..........
		....#.....
		.#......#.
		..........
		......#...
	`))

	assert.Equal(t, swappedUniverse, swapRowsColumns(universe))

	// Swapping again should get us back to where we started
	assert.Equal(t, universe, swapRowsColumns(swapRowsColumns(universe)))
}

func TestFindGalaxyLocations(t *testing.T) {
	universe := parseUniverse(heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....
	`))
	expandedUniverse := expandUniverse(universe)

	galaxyLocations := []GalaxyLocation{}
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 4,
		y: 0,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 9,
		y: 1,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 0,
		y: 2,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 8,
		y: 5,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 1,
		y: 6,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 12,
		y: 7,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 9,
		y: 10,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 0,
		y: 11,
	})
	galaxyLocations = append(galaxyLocations, GalaxyLocation{
		x: 5,
		y: 11,
	})

	assert.Equal(t, galaxyLocations, findGalaxyLocations(expandedUniverse))
}

func TestGetTaxiCabDistance(t *testing.T) {
	assert.Equal(t, 0, getTaxiCabDistance(0, 0, 0, 0))
	assert.Equal(t, 1, getTaxiCabDistance(1, 0, 0, 0))

	// Galaxies 5 and 9
	assert.Equal(t, 9, getTaxiCabDistance(1, 6, 5, 11))

	// Galaxies 1 and 7
	assert.Equal(t, 15, getTaxiCabDistance(4, 0, 9, 10))

	// Galaxies 8 and 9
	assert.Equal(t, 5, getTaxiCabDistance(0, 11, 5, 11))
}

func TestPairGalaxyLocations(t *testing.T) {
	universe := parseUniverse(heredoc.Doc(`
		...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....
	`))
	expandedUniverse := expandUniverse(universe)
	locations := findGalaxyLocations(expandedUniverse)
	pairs := pairGalaxyLocations(locations)

	assert.Equal(t, 36, len(pairs))
}
