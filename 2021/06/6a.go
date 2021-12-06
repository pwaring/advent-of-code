package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type Fish struct {
		timer int
	}

	const newFishDefaultTimer = 8

	scanner := bufio.NewScanner(os.Stdin)
	fish := []Fish{}
	maxDay := 80

	if scanner.Scan() {
		fishStrs := strings.Split(scanner.Text(), ",")

		for f := range fishStrs {
			tmpTimer, _ := strconv.Atoi(fishStrs[f])
			fish = append(fish, Fish{
				timer: tmpTimer,
			})
		}
	}

	for currentDay := 1; currentDay <= maxDay; currentDay++ {
		newFishCount := 0

		for f := range fish {
			if fish[f].timer == 0 {
				fish[f].timer = 6
				newFishCount++
			} else {
				fish[f].timer--
			}
		}

		for nf := 1; nf <= newFishCount; nf++ {
			fish = append(fish, Fish{
				timer: newFishDefaultTimer,
			})
		}
	}

	fmt.Println("After ", maxDay, " days there will be ", len(fish), " fish")
}
