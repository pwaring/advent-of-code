package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRaces(t *testing.T) {
	raceStr := "Time:      7  15   30\n"
	raceStr += "Distance:  9  40  200\n"

	race := Race{
		time:           71530,
		recordDistance: 940200,
	}

	assert.Equal(t, race, getRace(raceStr))
}

func TestGetRaceWinningOptionsCount(t *testing.T) {
	assert.Equal(t, 71503, getRaceWinningOptionsCount(Race{
		time:           71530,
		recordDistance: 940200,
	}))
}
