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
		Height      int
		scenicScore int
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
					Height: height,
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
			// If the tree is at the edge of the grid, it will always have
			// a scenic score of zero
			if row == 0 || row == gridRows-1 || column == 0 || column == gridColumns-1 {
				grid[row][column].scenicScore = 0
			} else {
				// Tree is not at the edge of the grid, so we need to check all trees
				// to the left, right, up, and down, moving outwards from the tree
				leftViewingDistance := 0
				rightViewingDistance := 0
				upViewingDistance := 0
				downViewingDistance := 0

				// Left
				for leftRow, leftHigherTreeFound := row-1, false; leftRow >= 0 && !leftHigherTreeFound; leftRow-- {
					if grid[leftRow][column].Height >= grid[row][column].Height {
						leftHigherTreeFound = true
					}

					leftViewingDistance++
				}

				// Right
				for rightRow, rightHigherTreeFound := row+1, false; rightRow < gridRows && !rightHigherTreeFound; rightRow++ {
					if grid[rightRow][column].Height >= grid[row][column].Height {
						rightHigherTreeFound = true
					}

					rightViewingDistance++
				}

				// Up
				for upColumn, upHigherTreeFound := column-1, false; upColumn >= 0 && !upHigherTreeFound; upColumn-- {
					if grid[row][upColumn].Height >= grid[row][column].Height {
						upHigherTreeFound = true
					}

					upViewingDistance++
				}

				// Down
				for downColumn, downHigherTreeFound := column+1, false; downColumn < gridColumns && !downHigherTreeFound; downColumn++ {
					if grid[row][downColumn].Height >= grid[row][column].Height {
						downHigherTreeFound = true
					}

					downViewingDistance++
				}

				// Calculate the scenic score as the product of viewing distances
				grid[row][column].scenicScore = leftViewingDistance * rightViewingDistance * upViewingDistance * downViewingDistance
			}
		}
	}

	// Scan the grid and find the highest scenic score
	// Start at zero since we know at least one tree will have this score
	highestScenicScore := 0

	for row := range grid {
		for column := range grid[row] {
			if grid[row][column].scenicScore > highestScenicScore {
				highestScenicScore = grid[row][column].scenicScore
			}
		}
	}

	fmt.Println(highestScenicScore)
}
