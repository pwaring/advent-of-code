package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSequence(t *testing.T) {
	assert.Equal(t, Sequence{
		readings: []int{0, 3, 6, 9, 12, 15},
	}, parseSequence("0 3 6 9 12 15"))

	assert.Equal(t, Sequence{
		readings: []int{1, 3, 6, 10, 15, 21},
	}, parseSequence("1 3 6 10 15 21"))

	assert.Equal(t, Sequence{
		readings: []int{10, 13, 16, 21, 30, 45},
	}, parseSequence("10 13 16 21 30 45"))
}

func TestIsZero(t *testing.T) {
	assert.Equal(t, false, isZero(Sequence{
		readings: []int{0, 3, 6, 9, 12, 15},
	}))

	assert.Equal(t, false, isZero(Sequence{
		readings: []int{1, 3, 6, 10, 15, 21},
	}))

	assert.Equal(t, false, isZero(Sequence{
		readings: []int{10, 13, 16, 21, 30, 45},
	}))

	assert.Equal(t, false, isZero(Sequence{
		readings: []int{0, 0, 0, 0, 0, 1},
	}))

	assert.Equal(t, false, isZero(Sequence{
		readings: []int{1, 0, 0, 0, 0},
	}))

	assert.Equal(t, true, isZero(Sequence{
		readings: []int{},
	}))

	assert.Equal(t, true, isZero(Sequence{
		readings: []int{0},
	}))

	assert.Equal(t, true, isZero(Sequence{
		readings: []int{0, 0, 0, 0, 0},
	}))
}

func TestExtrapolateSequence(t *testing.T) {
	assert.Equal(t, []Sequence{
		{
			readings: []int{0, 3, 6, 9, 12, 15},
		},
		{
			readings: []int{3, 3, 3, 3, 3},
		},
		{
			readings: []int{0, 0, 0, 0},
		},
	}, extrapolateSequence(parseSequence("0 3 6 9 12 15")))
}

func TestPopulateEstimates(t *testing.T) {
	assert.Equal(t, []Sequence{
		{
			readings: []int{0, 3, 6, 9, 12, 15, 18},
		},
		{
			readings: []int{3, 3, 3, 3, 3, 3},
		},
		{
			readings: []int{0, 0, 0, 0, 0},
		},
	}, populateEstimates(extrapolateSequence(parseSequence("0 3 6 9 12 15"))))

	assert.Equal(t, []Sequence{
		{
			readings: []int{1, 3, 6, 10, 15, 21, 28},
		},
		{
			readings: []int{2, 3, 4, 5, 6, 7},
		},
		{
			readings: []int{1, 1, 1, 1, 1},
		},
		{
			readings: []int{0, 0, 0, 0},
		},
	}, populateEstimates(extrapolateSequence(parseSequence("1 3 6 10 15 21"))))

	assert.Equal(t, []Sequence{
		{
			readings: []int{10, 13, 16, 21, 30, 45, 68},
		},
		{
			readings: []int{3, 3, 5, 9, 15, 23},
		},
		{
			readings: []int{0, 2, 4, 6, 8},
		},
		{
			readings: []int{2, 2, 2, 2},
		},
		{
			readings: []int{0, 0, 0},
		},
	}, populateEstimates(extrapolateSequence(parseSequence("10 13 16 21 30 45"))))
}
