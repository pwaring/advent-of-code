package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRaces(t *testing.T) {
	racesStr := "Time:      7  15   30\n"
	racesStr += "Distance:  9  40  200\n"

	races := []Race{}
	races = append(races, Race{
		time:           7,
		recordDistance: 9,
	})
	races = append(races, Race{
		time:           15,
		recordDistance: 40,
	})
	races = append(races, Race{
		time:           30,
		recordDistance: 200,
	})

	assert.Equal(t, races, getRaces(racesStr))
}

func TestGetRaceWinningOptionsCount(t *testing.T) {
	assert.Equal(t, 4, getRaceWinningOptionsCount(Race{
		time:           7,
		recordDistance: 9,
	}))

	assert.Equal(t, 8, getRaceWinningOptionsCount(Race{
		time:           15,
		recordDistance: 40,
	}))

	assert.Equal(t, 9, getRaceWinningOptionsCount(Race{
		time:           30,
		recordDistance: 200,
	}))
}
