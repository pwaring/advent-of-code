package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id                 int
	winningNumbers     []int
	playerNumbers      []int
	playerWinningCount int
	points             int
}

func getCard(str string) Card {
	card := Card{}

	cardParts := strings.Split(str, ":")
	card.id = getCardId(cardParts[0])

	numberParts := strings.Split(cardParts[1], "|")
	card.winningNumbers = getNumbers(numberParts[0])
	card.playerNumbers = getNumbers(numberParts[1])

	return card
}

func getCardId(str string) int {
	// Card ID strings are: 'Card ' followed by the ID
	// Example: Card 1 = ID of 1
	// So the ID is everything after character 5
	// Assumption: ASCII input so char = byte
	id, _ := strconv.Atoi(str[5:])
	return id
}

func getNumbers(str string) []int {
	numbers := []int{}

	str = strings.TrimSpace(str)
	strNumbers := strings.Fields(str)
	for sn := range strNumbers {
		number, _ := strconv.Atoi(strNumbers[sn])
		numbers = append(numbers, number)
	}

	return numbers
}

func getPlayerWinningCount(card Card) int {
	playerWinningCount := 0

	for pn := range card.playerNumbers {
		if slices.Contains(card.winningNumbers, card.playerNumbers[pn]) {
			playerWinningCount++
		}
	}

	return playerWinningCount
}

func getPoints(playerWinningCount int) int {
	points := 0

	for pwc := 1; pwc <= playerWinningCount; pwc++ {
		if pwc == 1 {
			points = 1
		} else {
			points = points * 2
		}
	}

	return points
}

func main() {
	cards := []Card{}
	scanner := bufio.NewScanner(os.Stdin)

	// Read in each line and convert to card structure
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			cards = append(cards, getCard(line))
		}
	}

	// Populate points
	for c := range cards {
		cards[c].playerWinningCount = getPlayerWinningCount(cards[c])
		cards[c].points = getPoints(cards[c].playerWinningCount)
	}

	// Calculate total points
	totalPoints := 0
	for c := range cards {
		totalPoints += cards[c].points
	}

	fmt.Println(totalPoints)
}
