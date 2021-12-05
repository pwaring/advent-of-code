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
		start Point
		end   Point
		all   []Point
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
	// ASSUMPTION: diagonal lines will be 45 degrees, i.e. x and y will move in step
	for lineIndex := range lines {
		currentLine := lines[lineIndex]

		xStep := 0
		yStep := 0

		if currentLine.start.x < currentLine.end.x {
			xStep = 1
		} else if currentLine.start.x > currentLine.end.x {
			xStep = -1
		}

		if currentLine.start.y < currentLine.end.y {
			yStep = 1
		} else if currentLine.start.y > currentLine.end.y {
			yStep = -1
		}

		// Add all the points from the start until (but not including) the end
		for currentX, currentY := currentLine.start.x, currentLine.start.y; currentX != currentLine.end.x || currentY != currentLine.end.y; currentX, currentY = currentX+xStep, currentY+yStep {
			lines[lineIndex].all = append(lines[lineIndex].all, Point{
				x: currentX,
				y: currentY,
			})
		}

		// Add the end point
		lines[lineIndex].all = append(lines[lineIndex].all, Point{
			x: currentLine.end.x,
			y: currentLine.end.y,
		})
	}

	// Calculate all overlaps, but only on horizontal and vertical lines
	for lineIndex := range lines {
		currentLine := lines[lineIndex]

		for pointIndex := range currentLine.all {
			overlaps[currentLine.all[pointIndex]]++
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
