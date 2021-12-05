package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const boardColumnCount int = 5
	const boardRowCount = 5

	type BoardItem struct {
		value  int
		marked bool
	}

	type Board struct {
		won         bool
		winningDraw int
		data        [boardColumnCount][boardRowCount]BoardItem
	}

	draws := []int{}
	boards := []Board{}
	currentBoard := 0
	nextBoard := 0
	currentBoardLine := 0
	currentBoardInitialised := false

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		drawStrs := strings.Split(scanner.Text(), ",")

		for d := range drawStrs {
			number, _ := strconv.Atoi(drawStrs[d])
			draws = append(draws, number)
		}
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			if !currentBoardInitialised {
				boards = append(boards, Board{})
				currentBoardInitialised = true
				currentBoardLine = 0
				nextBoard++
			}

			// Get every character in this line and convert to an integer
			numbers := strings.Fields(line)

			for n := range numbers {
				boards[currentBoard].data[currentBoardLine][n].value, _ = strconv.Atoi(numbers[n])
			}

			currentBoardLine++
		} else {
			// We may enter this part multiple times if multiple empty lines separate
			// each board in the input, but as these are assignments to values which
			// do not change, they are effectively idempotent
			currentBoardInitialised = false
			currentBoard = nextBoard
		}
	}

	lastBingoBoard := Board{}

	for drawIndex := 0; drawIndex < len(draws); drawIndex++ {
		// Mark all the numbers which match the current draw
		for boardIndex := range boards {
			for columnIndex := range boards[boardIndex].data {
				for rowIndex := range boards[boardIndex].data[columnIndex] {
					if boards[boardIndex].data[columnIndex][rowIndex].value == draws[drawIndex] {
						boards[boardIndex].data[columnIndex][rowIndex].marked = true
					}
				}
			}
		}

		// Find the first board with either all columns marked or all rows marked
		// Such a board may not exist on this draw
		for boardIndex := range boards {
			// Only bother to check this board if it has not already been won, as we
			// want to find the last board for which this is the case
			if !boards[boardIndex].won {
				for columnIndex := 0; columnIndex < boardColumnCount; columnIndex++ {
					markedColumnCount := 0

					for rowIndex := 0; rowIndex < boardRowCount; rowIndex++ {
						if boards[boardIndex].data[columnIndex][rowIndex].marked {
							markedColumnCount++
						}
					}

					if markedColumnCount == boardColumnCount {
						boards[boardIndex].won = true
						boards[boardIndex].winningDraw = draws[drawIndex]
						lastBingoBoard = boards[boardIndex]
					}
				}

				// Check all rows for marks
				for rowIndex := 0; rowIndex < boardRowCount; rowIndex++ {
					markedRowCount := 0

					for columnIndex := 0; columnIndex < boardColumnCount; columnIndex++ {
						if boards[boardIndex].data[columnIndex][rowIndex].marked {
							markedRowCount++
						}
					}

					if markedRowCount == boardColumnCount {
						boards[boardIndex].won = true
						boards[boardIndex].winningDraw = draws[drawIndex]
						lastBingoBoard = boards[boardIndex]
					}
				}
			}
		}
	}

	if lastBingoBoard.won {
		sumUnmarked := 0

		for columnIndex := range lastBingoBoard.data {
			for rowIndex := range lastBingoBoard.data[columnIndex] {
				if !lastBingoBoard.data[columnIndex][rowIndex].marked {
					sumUnmarked += lastBingoBoard.data[columnIndex][rowIndex].value
				}
			}
		}

		score := sumUnmarked * lastBingoBoard.winningDraw

		fmt.Println("Final score: ", score)
	} else {
		fmt.Println("No winning board found")
	}
}
