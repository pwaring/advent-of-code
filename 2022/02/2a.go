package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		Rock = iota
		Paper
		Scissors
	)

	type Round struct {
		elfHand     int
		playerHand  int
		playerScore int
	}

	rounds := []Round{}

	scanner := bufio.NewScanner(os.Stdin)

	// Read in the data and convert into structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		hands := strings.Fields(line)
		elfHandString := hands[0]
		playerHandString := hands[1]

		round := Round{}

		if elfHandString == "A" {
			round.elfHand = Rock
		} else if elfHandString == "B" {
			round.elfHand = Paper
		} else if elfHandString == "C" {
			round.elfHand = Scissors
		}

		if playerHandString == "X" {
			round.playerHand = Rock
		} else if playerHandString == "Y" {
			round.playerHand = Paper
		} else if playerHandString == "Z" {
			round.playerHand = Scissors
		}

		round.playerScore = 0

		if round.playerHand == Rock {
			round.playerScore += 1
		} else if round.playerHand == Paper {
			round.playerScore += 2
		} else if round.playerHand == Scissors {
			round.playerScore += 3
		}

		if round.elfHand == round.playerHand {
			// Draw
			round.playerScore += 3
		} else if round.elfHand == Rock && round.playerHand == Paper {
			// Win: Paper wraps Rock
			round.playerScore += 6
		} else if round.elfHand == Scissors && round.playerHand == Rock {
			// Win: Rock breaks Scissors
			round.playerScore += 6
		} else if round.elfHand == Paper && round.playerHand == Scissors {
			// Win: Scissors cut Paper
			round.playerScore += 6
		}

		rounds = append(rounds, round)
	}

	// Calculate total score
	totalScore := 0

	for r := range rounds {
		totalScore += rounds[r].playerScore
	}

	fmt.Println(totalScore)
}
