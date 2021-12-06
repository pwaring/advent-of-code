package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const newFishDefaultTimer = 8

	scanner := bufio.NewScanner(os.Stdin)
	fish := [9]int{}
	maxDay := 256

	if scanner.Scan() {
		fishStrs := strings.Split(scanner.Text(), ",")

		for f := range fishStrs {
			timer, _ := strconv.Atoi(fishStrs[f])
			fish[timer]++
		}
	}

	for currentDay := 1; currentDay <= maxDay; currentDay++ {
		// Could probably do this without a temporary array, but this is simple
		// and has minimal overhead given the array size
		tmpFish := [9]int{}

		for f := 1; f < len(fish); f++ {
			tmpFish[f-1] = fish[f]
		}

		// All fish at 0 timer create new fish with a timer of 8
		// and the birthing fish has its timer reset to 6
		tmpFish[8] += fish[0]
		tmpFish[6] += fish[0]

		fish = tmpFish
	}

	fishCount := 0

	for f := range fish {
		fishCount += fish[f]
	}

	fmt.Println("After ", maxDay, " days there will be ", fishCount, " fish")
}
