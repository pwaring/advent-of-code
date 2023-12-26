package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSpringsRow(t *testing.T) {
	assert.Equal(t, SpringsRow{
		springs: []int{
			Unknown,
			Unknown,
			Unknown,
			Operational,
			Damaged,
			Damaged,
			Damaged,
		},
		damagedGroups: []int{
			1,
			1,
			3,
		},
	}, parseSpringsRow("???.### 1,1,3"))

	assert.Equal(t, SpringsRow{
		springs: []int{
			Operational,
			Unknown,
			Unknown,
			Operational,
			Operational,
			Unknown,
			Unknown,
			Operational,
			Operational,
			Operational,
			Unknown,
			Damaged,
			Damaged,
			Operational,
		},
		damagedGroups: []int{
			1,
			1,
			3,
		},
	}, parseSpringsRow(".??..??...?##. 1,1,3"))

	assert.Equal(t, SpringsRow{
		springs: []int{
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
		},
		damagedGroups: []int{
			1,
			3,
			1,
			6,
		},
	}, parseSpringsRow("?#?#?#?#?#?#?#? 1,3,1,6"))

	assert.Equal(t, SpringsRow{
		springs: []int{
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Operational,
			Damaged,
			Operational,
			Operational,
			Operational,
			Damaged,
			Operational,
			Operational,
			Operational,
		},
		damagedGroups: []int{
			4,
			1,
			1,
		},
	}, parseSpringsRow("????.#...#... 4,1,1"))

	assert.Equal(t, SpringsRow{
		springs: []int{
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Operational,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Operational,
			Operational,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Operational,
		},
		damagedGroups: []int{
			1,
			6,
			5,
		},
	}, parseSpringsRow("????.######..#####. 1,6,5"))

	assert.Equal(t, SpringsRow{
		springs: []int{
			Unknown,
			Damaged,
			Damaged,
			Damaged,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
		},
		damagedGroups: []int{
			3,
			2,
			1,
		},
	}, parseSpringsRow("?###???????? 3,2,1"))
}

func TestDrawSpringsRow(t *testing.T) {
	assert.Equal(t, "???.### 1,1,3", drawSpringsRow(SpringsRow{
		springs: []int{
			Unknown,
			Unknown,
			Unknown,
			Operational,
			Damaged,
			Damaged,
			Damaged,
		},
		damagedGroups: []int{
			1,
			1,
			3,
		},
	}))

	assert.Equal(t, ".??..??...?##. 1,1,3", drawSpringsRow(SpringsRow{
		springs: []int{
			Operational,
			Unknown,
			Unknown,
			Operational,
			Operational,
			Unknown,
			Unknown,
			Operational,
			Operational,
			Operational,
			Unknown,
			Damaged,
			Damaged,
			Operational,
		},
		damagedGroups: []int{
			1,
			1,
			3,
		},
	}))

	assert.Equal(t, "?#?#?#?#?#?#?#? 1,3,1,6", drawSpringsRow(SpringsRow{
		springs: []int{
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
			Damaged,
			Unknown,
		},
		damagedGroups: []int{
			1,
			3,
			1,
			6,
		},
	}))

	assert.Equal(t, "????.#...#... 4,1,1", drawSpringsRow(SpringsRow{
		springs: []int{
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Operational,
			Damaged,
			Operational,
			Operational,
			Operational,
			Damaged,
			Operational,
			Operational,
			Operational,
		},
		damagedGroups: []int{
			4,
			1,
			1,
		},
	}))

	assert.Equal(t, "????.######..#####. 1,6,5", drawSpringsRow(SpringsRow{
		springs: []int{
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Operational,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Operational,
			Operational,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Damaged,
			Operational,
		},
		damagedGroups: []int{
			1,
			6,
			5,
		},
	}))

	assert.Equal(t, "?###???????? 3,2,1", drawSpringsRow(SpringsRow{
		springs: []int{
			Unknown,
			Damaged,
			Damaged,
			Damaged,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
			Unknown,
		},
		damagedGroups: []int{
			3,
			2,
			1,
		},
	}))
}

func TestCalculateDamagedGroups(t *testing.T) {
	assert.Equal(t, []int{3}, calculateDamagedGroups([]int{
		Operational,
		Operational,
		Operational,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))

	assert.Equal(t, []int{1, 3}, calculateDamagedGroups([]int{
		Operational,
		Operational,
		Damaged,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))

	assert.Equal(t, []int{1, 3}, calculateDamagedGroups([]int{
		Operational,
		Damaged,
		Operational,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))

	assert.Equal(t, []int{2, 3}, calculateDamagedGroups([]int{
		Operational,
		Damaged,
		Damaged,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))

	assert.Equal(t, []int{1, 3}, calculateDamagedGroups([]int{
		Damaged,
		Operational,
		Operational,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))

	assert.Equal(t, []int{1, 1, 3}, calculateDamagedGroups([]int{
		Damaged,
		Operational,
		Damaged,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))

	assert.Equal(t, []int{2, 3}, calculateDamagedGroups([]int{
		Damaged,
		Damaged,
		Operational,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))

	assert.Equal(t, []int{3, 3}, calculateDamagedGroups([]int{
		Damaged,
		Damaged,
		Damaged,
		Operational,
		Damaged,
		Damaged,
		Damaged,
	}))
}

func TestGetCombinations(t *testing.T) {
	choices := []int{Operational, Damaged}

	assert.Equal(t, [][]int{
		{1},
	}, getCombinations([]int{1}, 1))

	assert.Equal(t, [][]int{
		{1, 1},
	}, getCombinations([]int{1}, 2))

	assert.Equal(t, [][]int{
		{Operational, Operational, Operational},
		{Operational, Operational, Damaged},
		{Operational, Damaged, Operational},
		{Operational, Damaged, Damaged},
		{Damaged, Operational, Operational},
		{Damaged, Operational, Damaged},
		{Damaged, Damaged, Operational},
		{Damaged, Damaged, Damaged},
	}, getCombinations(choices, 3))

	assert.Len(t, getCombinations(choices, 5), 32)
	assert.Len(t, getCombinations(choices, 8), 256)
}

func TestCountMatchingArrangements(t *testing.T) {
	choices := []int{Operational, Damaged}

	assert.Equal(t, 1, countMatchingArrangements(
		parseSpringsRow("???.### 1,1,3"),
		choices,
	))

	assert.Equal(t, 4, countMatchingArrangements(
		parseSpringsRow(".??..??...?##. 1,1,3"),
		choices,
	))

	assert.Equal(t, 1, countMatchingArrangements(
		parseSpringsRow("?#?#?#?#?#?#?#? 1,3,1,6"),
		choices,
	))

	assert.Equal(t, 1, countMatchingArrangements(
		parseSpringsRow("????.#...#... 4,1,1"),
		choices,
	))

	assert.Equal(t, 4, countMatchingArrangements(
		parseSpringsRow("????.######..#####. 1,6,5"),
		choices,
	))

	assert.Equal(t, 10, countMatchingArrangements(
		parseSpringsRow("?###???????? 3,2,1"),
		choices,
	))
}

func TestPopulatePlaceholders(t *testing.T) {
	assert.Equal(t, []int{
		Damaged,
		Operational,
		Damaged,
		Operational,
		Damaged,
		Damaged,
	}, populatePlaceholders([]int{
		Unknown,
		Unknown,
		Damaged,
		Operational,
		Damaged,
		Damaged,
	}, []int{
		Damaged,
		Operational,
	}))
}
