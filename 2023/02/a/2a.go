package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reveal struct {
	red   int
	blue  int
	green int
}

type Game struct {
	id      int
	reveals []Reveal
}

func getGameId(str string) int {
	// Game ID strings are: 'Game ' followed by the ID
	// So the ID is everything after character 5
	id, _ := strconv.Atoi(str[5:])
	return id
}

func getReveal(str string) Reveal {
	// A reveal consists of zero or more cube definitions
	// Cube definition is: 'N Colour', e.g. '2 red'
	reveal := Reveal{}

	revealParts := strings.Split(str, ",")

	for rp := range revealParts {
		currentRevealPart := strings.TrimSpace(revealParts[rp])
		cubeParts := strings.Fields(currentRevealPart)
		count, _ := strconv.Atoi(strings.TrimSpace(cubeParts[0]))
		colour := strings.TrimSpace(cubeParts[1])

		if colour == "red" {
			reveal.red = count
		} else if colour == "blue" {
			reveal.blue = count
		} else if colour == "green" {
			reveal.green = count
		}
	}

	return reveal
}

func getReveals(str string) []Reveal {
	reveals := make([]Reveal, 0)

	// A group of reveals is separated by semi-colons
	parts := strings.Split(str, ";")
	for p := range parts {
		reveals = append(reveals, getReveal(strings.TrimSpace(parts[p])))
	}

	return reveals
}

func main() {
	games := []Game{}
	scanner := bufio.NewScanner(os.Stdin)

	// Read each line and convert into data structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			game := Game{}
			lineParts := strings.Split(line, ":")

			game.id = getGameId(lineParts[0])
			game.reveals = getReveals(lineParts[1])

			games = append(games, game)
		}
	}

	originalBag := Reveal{
		red:   12,
		blue:  14,
		green: 13,
	}

	idSum := 0

	for g := range games {
		matchedReveals := 0
		totalReveals := len(games[g].reveals)

		for r := range games[g].reveals {
			currentReveal := games[g].reveals[r]

			if currentReveal.red <= originalBag.red && currentReveal.blue <= originalBag.blue && currentReveal.green <= originalBag.green {
				matchedReveals++
			}
		}

		if matchedReveals == totalReveals {
			idSum += games[g].id
		}
	}

	fmt.Println(idSum)
}
