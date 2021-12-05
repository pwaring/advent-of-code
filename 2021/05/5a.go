package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type Point struct {
		x int
		y int
	}

	type Line struct {
		start      Point
		end        Point
		horizontal bool
		vertical   bool
		all        []Point
	}

	const OVERLAP_TARGET = 2

	lines := []Line{}
	overlaps := make(map[Point]int)

	// Read input and convert to data structures
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " -> ")

		startPoint := Point{}
		endPoint := Point{}
		line := Line{}

		tmpPoint := strings.Split(input[0], ",")
		startPoint.x, _ = strconv.Atoi(tmpPoint[0])
		startPoint.y, _ = strconv.Atoi(tmpPoint[1])

		tmpPoint = strings.Split(input[1], ",")
		endPoint.x, _ = strconv.Atoi(tmpPoint[0])
		endPoint.y, _ = strconv.Atoi(tmpPoint[1])

		line.start = startPoint
		line.end = endPoint

		lines = append(lines, line)
	}

	// Calculate all the points each line goes through
	for lineIndex := range lines {
		currentLine := lines[lineIndex]

		// Vertical lines
		if currentLine.start.x == currentLine.end.x {
			yStep := 0

			if currentLine.start.y < currentLine.end.y {
				yStep = 1
			} else if currentLine.start.y > currentLine.end.y {
				yStep = -1
			}

			for currentY := currentLine.start.y; currentY != currentLine.end.y; currentY += yStep {
				lines[lineIndex].all = append(lines[lineIndex].all, Point{
					x: currentLine.start.x,
					y: currentY,
				})
			}

			// Add the end point
			lines[lineIndex].all = append(lines[lineIndex].all, Point{
				x: currentLine.end.x,
				y: currentLine.end.y,
			})

			// Mark as vertical
			lines[lineIndex].vertical = true
		}

		// Horizontal lines
		if currentLine.start.y == currentLine.end.y {
			xStep := 0

			if currentLine.start.x < currentLine.end.x {
				xStep = 1
			} else if currentLine.start.x > currentLine.end.x {
				xStep = -1
			}

			for currentX := currentLine.start.x; currentX != currentLine.end.x; currentX += xStep {
				lines[lineIndex].all = append(lines[lineIndex].all, Point{
					x: currentX,
					y: currentLine.start.y,
				})
			}

			// Add the end point
			lines[lineIndex].all = append(lines[lineIndex].all, Point{
				x: currentLine.end.x,
				y: currentLine.end.y,
			})

			// Mark as horizontal
			lines[lineIndex].horizontal = true
		}
	}

	// Calculate all overlaps, but only on horizontal and vertical lines
	for lineIndex := range lines {
		currentLine := lines[lineIndex]

		if currentLine.vertical || currentLine.horizontal {
			for pointIndex := range currentLine.all {
				overlaps[currentLine.all[pointIndex]]++
			}
		}
	}

	// Count number of overlaps at or above target
	overlapTargetCount := 0

	for _, overlapCount := range overlaps {
		if overlapCount >= OVERLAP_TARGET {
			overlapTargetCount++
		}
	}

	fmt.Println("Points with at least ", OVERLAP_TARGET, " overlapping lines: ", overlapTargetCount)
}
