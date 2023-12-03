package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EngineCell struct {
	content  string
	isDigit  bool
	isSymbol bool
}

func getPartNumbers(engine [][]EngineCell) []int {
	partNumbers := []int{}

	for r := range engine {
		engineRow := engine[r]
		possiblePartNumber := ""
		adjacentSymbolCount := 0

		for c := range engineRow {
			if engineRow[c].isDigit {
				possiblePartNumber += engineRow[c].content
				adjacentCells := getAdjacentCells(r, c, engine)

				for ac := range adjacentCells {
					if adjacentCells[ac].isSymbol {
						adjacentSymbolCount++
					}
				}

				// Edge case: we're at the end of the row
				if c == len(engineRow)-1 && adjacentSymbolCount > 0 {
					partNumber, _ := strconv.Atoi(possiblePartNumber)
					partNumbers = append(partNumbers, partNumber)
				}
			} else {
				// No more digits in the current part number, so check
				// whether it has adjacent symbols
				if adjacentSymbolCount > 0 {
					partNumber, _ := strconv.Atoi(possiblePartNumber)
					partNumbers = append(partNumbers, partNumber)
				}

				// Reset the possible part number and adjacent symbol
				// count in preparation for the next number
				possiblePartNumber = ""
				adjacentSymbolCount = 0
			}
		}
	}

	return partNumbers
}

func getAdjacentCells(row int, column int, engine [][]EngineCell) []EngineCell {
	engineCells := make([]EngineCell, 0)

	// Top Left
	if cellExists(row-1, column-1, engine) {
		engineCells = append(engineCells, engine[row-1][column-1])
	}

	// Top Middle
	if cellExists(row-1, column, engine) {
		engineCells = append(engineCells, engine[row-1][column])
	}

	// Top Right
	if cellExists(row-1, column+1, engine) {
		engineCells = append(engineCells, engine[row-1][column+1])
	}

	// Middle Left
	if cellExists(row, column-1, engine) {
		engineCells = append(engineCells, engine[row][column-1])
	}

	// Middle Right
	if cellExists(row, column+1, engine) {
		engineCells = append(engineCells, engine[row][column+1])
	}

	// Bottom Left
	if cellExists(row+1, column-1, engine) {
		engineCells = append(engineCells, engine[row+1][column-1])
	}

	// Bottom Middle
	if cellExists(row+1, column, engine) {
		engineCells = append(engineCells, engine[row+1][column])
	}

	// Bottom Right
	if cellExists(row+1, column+1, engine) {
		engineCells = append(engineCells, engine[row+1][column+1])
	}

	return engineCells
}

func cellExists(row int, column int, engine [][]EngineCell) bool {
	return row >= 0 && row < len(engine) && column >= 0 && column < len(engine[row])
}

func isDigit(str string) bool {
	if len(str) == 1 {
		_, err := strconv.Atoi(str)
		return err == nil
	}

	return false
}

func isSymbol(str string) bool {
	// A symbol is anything other than a digit or full stop
	return !isDigit(str) && str != "."
}

func rowToEngineCells(row string) []EngineCell {
	engineCells := make([]EngineCell, len(row))

	characters := strings.Split(row, "")

	for c := range characters {
		engineCells[c].content = characters[c]
		engineCells[c].isDigit = isDigit(characters[c])
		engineCells[c].isSymbol = isSymbol(characters[c])
	}

	return engineCells
}

func main() {
	engine := [][]EngineCell{}

	scanner := bufio.NewScanner(os.Stdin)

	// Read in each line and convert into cell structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			engine = append(engine, rowToEngineCells(line))
		}
	}

	partNumbers := getPartNumbers(engine)

	partNumberSum := 0

	for pn := range partNumbers {
		partNumberSum += partNumbers[pn]
	}

	fmt.Println(partNumberSum)
}
