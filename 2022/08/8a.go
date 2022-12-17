package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type Tree struct {
		Height  int
		Visible bool
	}

	grid := [][]Tree{}

	scanner := bufio.NewScanner(os.Stdin)

	// Read in each line and convert into grid structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Only process line if we have content - may have an empty line at the end of the file
		if len(line) > 0 {
			gridRow := []Tree{}
			trees := strings.Split(line, "")

			for t := range trees {
				height, _ := strconv.Atoi(trees[t])

				tree := Tree{
					Height:  height,
					Visible: false,
				}

				gridRow = append(gridRow, tree)
			}

			grid = append(grid, gridRow)
		}
	}

	gridRows := len(grid)
	gridColumns := len(grid[0])

	// Scan the grid and mark trees as visible
	for row := range grid {
		for column := range grid[row] {
			// If the tree is at the edge of the grid, it is always visible
			if row == 0 || row == gridRows-1 || column == 0 || column == gridColumns-1 {
				grid[row][column].Visible = true
			} else {
				// Tree is not at the edge of the grid, so we need to check all trees
				// to the left, right, up, and down
				// We can stop once a tree is visible from any direction

				// Left
				if !grid[row][column].Visible {
					leftTreeTotalCount := row
					leftTreeShorterCount := 0

					for leftRow := 0; leftRow < row; leftRow++ {
						if grid[leftRow][column].Height < grid[row][column].Height {
							leftTreeShorterCount++
						}
					}

					if leftTreeShorterCount == leftTreeTotalCount {
						grid[row][column].Visible = true
					}
				}

				// Right
				if !grid[row][column].Visible {
					rightTreeTotalCount := (gridRows - 1) - row
					rightTreeShorterCount := 0

					for rightRow := gridRows - 1; rightRow > row; rightRow-- {
						if grid[rightRow][column].Height < grid[row][column].Height {
							rightTreeShorterCount++
						}
					}

					if rightTreeShorterCount == rightTreeTotalCount {
						grid[row][column].Visible = true
					}
				}

				// Up
				if !grid[row][column].Visible {
					upTreeTotalCount := column
					upTreeShorterCount := 0

					for upColumn := 0; upColumn < column; upColumn++ {
						if grid[row][upColumn].Height < grid[row][column].Height {
							upTreeShorterCount++
						}
					}

					if upTreeShorterCount == upTreeTotalCount {
						grid[row][column].Visible = true
					}
				}

				// Down
				if !grid[row][column].Visible {
					downTreeTotalCount := (gridColumns - 1) - column
					downTreeShorterCount := 0

					for downColumn := gridColumns - 1; downColumn > column; downColumn-- {
						if grid[row][downColumn].Height < grid[row][column].Height {
							downTreeShorterCount++
						}

						if downTreeShorterCount == downTreeTotalCount {
							grid[row][column].Visible = true
						}
					}
				}
			}
		}
	}

	// Scan the grid and count visible trees
	// We could have done this in the previous step, but we are keeping the
	// two operations (mark and count) separate in case part 2 requires
	// different processing
	visibleTreeCount := 0

	for row := range grid {
		for column := range grid[row] {
			if grid[row][column].Visible {
				visibleTreeCount++
			}
		}
	}

	fmt.Println(visibleTreeCount)
}
