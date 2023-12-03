package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGameId(t *testing.T) {
	assert.Equal(t, 1, getGameId("Game 1"))
	assert.Equal(t, 2, getGameId("Game 2"))
	assert.Equal(t, 12, getGameId("Game 12"))
	assert.Equal(t, 123, getGameId("Game 123"))
	assert.Equal(t, 5678, getGameId("Game 5678"))
}

func TestGetReveal(t *testing.T) {
	reveal := Reveal{}

	reveal.red = 4
	reveal.blue = 3
	reveal.green = 0
	assert.Equal(t, reveal, getReveal("3 blue, 4 red"))

	reveal.red = 1
	reveal.blue = 6
	reveal.green = 2
	assert.Equal(t, reveal, getReveal("1 red, 2 green, 6 blue"))

	reveal.red = 1
	reveal.blue = 6
	reveal.green = 2
	assert.Equal(t, reveal, getReveal("1 red, 2 green, 6 blue"))

	reveal.red = 0
	reveal.blue = 1
	reveal.green = 2
	assert.Equal(t, reveal, getReveal("1 blue, 2 green"))

	reveal.red = 1
	reveal.blue = 4
	reveal.green = 3
	assert.Equal(t, reveal, getReveal("3 green, 4 blue, 1 red"))

	reveal.red = 0
	reveal.blue = 1
	reveal.green = 1
	assert.Equal(t, reveal, getReveal("1 green, 1 blue"))

	reveal.red = 20
	reveal.blue = 6
	reveal.green = 8
	assert.Equal(t, reveal, getReveal("8 green, 6 blue, 20 red"))

	reveal.red = 4
	reveal.blue = 5
	reveal.green = 13
	assert.Equal(t, reveal, getReveal("5 blue, 4 red, 13 green"))

	reveal.red = 1
	reveal.blue = 0
	reveal.green = 5
	assert.Equal(t, reveal, getReveal("5 green, 1 red"))

	reveal.red = 3
	reveal.blue = 6
	reveal.green = 1
	assert.Equal(t, reveal, getReveal("1 green, 3 red, 6 blue"))

	reveal.red = 6
	reveal.blue = 0
	reveal.green = 3
	assert.Equal(t, reveal, getReveal("3 green, 6 red"))

	reveal.red = 14
	reveal.blue = 15
	reveal.green = 3
	assert.Equal(t, reveal, getReveal("3 green, 15 blue, 14 red"))

	reveal.red = 6
	reveal.blue = 1
	reveal.green = 3
	assert.Equal(t, reveal, getReveal("6 red, 1 blue, 3 green"))
}

func TestGetReveals(t *testing.T) {
	reveals := make([]Reveal, 3)
	reveals[0].red = 4
	reveals[0].blue = 3
	reveals[0].green = 0
	reveals[1].red = 1
	reveals[1].blue = 6
	reveals[1].green = 2
	reveals[2].red = 0
	reveals[2].blue = 0
	reveals[2].green = 2
	assert.Equal(t, reveals, getReveals("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
}

func TestGetMinimumReveal(t *testing.T) {
	reveal := Reveal{
		red:   4,
		blue:  6,
		green: 2,
	}
	assert.Equal(t, reveal, getMinimumReveal(getReveals("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")))

	reveal = Reveal{
		red:   1,
		blue:  4,
		green: 3,
	}
	assert.Equal(t, reveal, getMinimumReveal(getReveals("1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")))

	reveal = Reveal{
		red:   20,
		blue:  6,
		green: 13,
	}
	assert.Equal(t, reveal, getMinimumReveal(getReveals("8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")))

	reveal = Reveal{
		red:   14,
		blue:  15,
		green: 3,
	}
	assert.Equal(t, reveal, getMinimumReveal(getReveals("1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")))

	reveal = Reveal{
		red:   6,
		blue:  2,
		green: 3,
	}
	assert.Equal(t, reveal, getMinimumReveal(getReveals("6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")))
}
