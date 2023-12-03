package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.Equal(t, true, isDigit("0"))
	assert.Equal(t, true, isDigit("1"))
	assert.Equal(t, true, isDigit("2"))
	assert.Equal(t, true, isDigit("3"))
	assert.Equal(t, true, isDigit("4"))
	assert.Equal(t, true, isDigit("5"))
	assert.Equal(t, true, isDigit("6"))
	assert.Equal(t, true, isDigit("7"))
	assert.Equal(t, true, isDigit("8"))
	assert.Equal(t, true, isDigit("9"))
	assert.Equal(t, false, isDigit("*"))
	assert.Equal(t, false, isDigit("$"))
	assert.Equal(t, false, isDigit("."))
	assert.Equal(t, false, isDigit("#"))
	assert.Equal(t, false, isDigit("+"))
	assert.Equal(t, false, isDigit("/"))
	assert.Equal(t, false, isDigit("%"))
	assert.Equal(t, false, isDigit("-"))
	assert.Equal(t, false, isDigit("@"))
	assert.Equal(t, false, isDigit("="))
	assert.Equal(t, false, isDigit("-1"))
}

func TestIsSymbol(t *testing.T) {
	assert.Equal(t, false, isSymbol("0"))
	assert.Equal(t, false, isSymbol("1"))
	assert.Equal(t, false, isSymbol("2"))
	assert.Equal(t, false, isSymbol("3"))
	assert.Equal(t, false, isSymbol("4"))
	assert.Equal(t, false, isSymbol("5"))
	assert.Equal(t, false, isSymbol("6"))
	assert.Equal(t, false, isSymbol("7"))
	assert.Equal(t, false, isSymbol("8"))
	assert.Equal(t, false, isSymbol("9"))
	assert.Equal(t, true, isSymbol("*"))
	assert.Equal(t, true, isSymbol("$"))
	assert.Equal(t, false, isSymbol("."))
	assert.Equal(t, true, isSymbol("#"))
	assert.Equal(t, true, isSymbol("+"))
	assert.Equal(t, true, isSymbol("/"))
	assert.Equal(t, true, isSymbol("%"))
	assert.Equal(t, true, isSymbol("-"))
	assert.Equal(t, true, isSymbol("@"))
	assert.Equal(t, true, isSymbol("="))
}

func TestRowToEngineCells(t *testing.T) {
	engineCells := make([]EngineCell, 10)
	engineCells[0].content = "4"
	engineCells[0].isDigit = true
	engineCells[0].isSymbol = false
	engineCells[1].content = "6"
	engineCells[1].isDigit = true
	engineCells[1].isSymbol = false
	engineCells[2].content = "7"
	engineCells[2].isDigit = true
	engineCells[2].isSymbol = false
	engineCells[3].content = "."
	engineCells[3].isDigit = false
	engineCells[3].isSymbol = false
	engineCells[4].content = "."
	engineCells[4].isDigit = false
	engineCells[4].isSymbol = false
	engineCells[5].content = "1"
	engineCells[5].isDigit = true
	engineCells[5].isSymbol = false
	engineCells[6].content = "1"
	engineCells[6].isDigit = true
	engineCells[6].isSymbol = false
	engineCells[7].content = "4"
	engineCells[7].isDigit = true
	engineCells[7].isSymbol = false
	engineCells[8].content = "."
	engineCells[8].isDigit = false
	engineCells[8].isSymbol = false
	engineCells[9].content = "."
	engineCells[9].isDigit = false
	engineCells[9].isSymbol = false

	assert.Equal(t, engineCells, rowToEngineCells("467..114.."))
}

func TestGetCellExists(t *testing.T) {
	engine := [][]EngineCell{}
	engine = append(engine, rowToEngineCells("467..114.."))
	engine = append(engine, rowToEngineCells("...*......"))
	engine = append(engine, rowToEngineCells("..35..633."))
	engine = append(engine, rowToEngineCells("......#..."))
	engine = append(engine, rowToEngineCells("617*......"))
	engine = append(engine, rowToEngineCells(".....+.58."))
	engine = append(engine, rowToEngineCells("..592....."))
	engine = append(engine, rowToEngineCells("......755."))
	engine = append(engine, rowToEngineCells("...$.*...."))
	engine = append(engine, rowToEngineCells(".664.598.."))

	assert.Equal(t, true, cellExists(0, 0, engine))
	assert.Equal(t, true, cellExists(9, 9, engine))
	assert.Equal(t, true, cellExists(3, 3, engine))
	assert.Equal(t, false, cellExists(-1, 0, engine))
	assert.Equal(t, false, cellExists(0, -1, engine))
	assert.Equal(t, false, cellExists(10, 0, engine))
	assert.Equal(t, false, cellExists(0, 10, engine))
}

func TestGetAdjacentCells(t *testing.T) {
	engine := [][]EngineCell{}
	engine = append(engine, rowToEngineCells("467..114.."))
	engine = append(engine, rowToEngineCells("...*......"))
	engine = append(engine, rowToEngineCells("..35..633."))
	engine = append(engine, rowToEngineCells("......#..."))
	engine = append(engine, rowToEngineCells("617*......"))
	engine = append(engine, rowToEngineCells(".....+.58."))
	engine = append(engine, rowToEngineCells("..592....."))
	engine = append(engine, rowToEngineCells("......755."))
	engine = append(engine, rowToEngineCells("...$.*...."))
	engine = append(engine, rowToEngineCells(".664.598.."))

	// 0, 0
	adjacentCells := []EngineCell{}
	adjacentCells = append(adjacentCells, engine[0][1])
	adjacentCells = append(adjacentCells, engine[1][0])
	adjacentCells = append(adjacentCells, engine[1][1])
	assert.Equal(t, adjacentCells, getAdjacentCells(0, 0, engine))

	// 5, 5
	adjacentCells = []EngineCell{}
	adjacentCells = append(adjacentCells, engine[4][4])
	adjacentCells = append(adjacentCells, engine[4][5])
	adjacentCells = append(adjacentCells, engine[4][6])
	adjacentCells = append(adjacentCells, engine[5][4])
	adjacentCells = append(adjacentCells, engine[5][6])
	adjacentCells = append(adjacentCells, engine[6][4])
	adjacentCells = append(adjacentCells, engine[6][5])
	adjacentCells = append(adjacentCells, engine[6][6])
	assert.Equal(t, adjacentCells, getAdjacentCells(5, 5, engine))

	// 9, 9
	adjacentCells = []EngineCell{}
	adjacentCells = append(adjacentCells, engine[8][8])
	adjacentCells = append(adjacentCells, engine[8][9])
	adjacentCells = append(adjacentCells, engine[9][8])
	assert.Equal(t, adjacentCells, getAdjacentCells(9, 9, engine))
}

func TestGetPartNumbers(t *testing.T) {
	engine := [][]EngineCell{}
	engine = append(engine, rowToEngineCells("467..114.."))
	engine = append(engine, rowToEngineCells("...*......"))
	engine = append(engine, rowToEngineCells("..35..633."))
	engine = append(engine, rowToEngineCells("......#..."))
	engine = append(engine, rowToEngineCells("617*......"))
	engine = append(engine, rowToEngineCells(".....+.58."))
	engine = append(engine, rowToEngineCells("..592....."))
	engine = append(engine, rowToEngineCells("......755."))
	engine = append(engine, rowToEngineCells("...$.*...."))
	engine = append(engine, rowToEngineCells(".664.598.."))

	partNumbers := []int{}
	partNumbers = append(partNumbers, 467)
	partNumbers = append(partNumbers, 35)
	partNumbers = append(partNumbers, 633)
	partNumbers = append(partNumbers, 617)
	partNumbers = append(partNumbers, 592)
	partNumbers = append(partNumbers, 755)
	partNumbers = append(partNumbers, 664)
	partNumbers = append(partNumbers, 598)
	assert.Equal(t, partNumbers, getPartNumbers(engine))

	partNumberSum := 0
	for pn := range partNumbers {
		partNumberSum += partNumbers[pn]
	}
	assert.Equal(t, 4361, partNumberSum)

	engine = [][]EngineCell{}
	engine = append(engine, rowToEngineCells("........"))
	engine = append(engine, rowToEngineCells(".24..4.."))
	engine = append(engine, rowToEngineCells("......*."))
	partNumbers = []int{}
	partNumbers = append(partNumbers, 4)
	assert.Equal(t, partNumbers, getPartNumbers(engine))

	engine = [][]EngineCell{}
	engine = append(engine, rowToEngineCells("........"))
	engine = append(engine, rowToEngineCells(".24..4.8"))
	engine = append(engine, rowToEngineCells("......*."))
	partNumbers = []int{}
	partNumbers = append(partNumbers, 4)
	partNumbers = append(partNumbers, 8)
	assert.Equal(t, partNumbers, getPartNumbers(engine))
}
