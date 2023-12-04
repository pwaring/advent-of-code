package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCardId(t *testing.T) {
	assert.Equal(t, 1, getCardId("Card 1"))
	assert.Equal(t, 2, getCardId("Card 2"))
	assert.Equal(t, 12, getCardId("Card 12"))
	assert.Equal(t, 123, getCardId("Card 123"))
	assert.Equal(t, 5678, getCardId("Card 5678"))
}

func TestGetNumbers(t *testing.T) {
	numbers := []int{}
	numbers = append(numbers, 41)
	numbers = append(numbers, 48)
	numbers = append(numbers, 83)
	numbers = append(numbers, 86)
	numbers = append(numbers, 17)
	assert.Equal(t, numbers, getNumbers("41 48 83 86 17"))

	numbers = []int{}
	numbers = append(numbers, 83)
	numbers = append(numbers, 86)
	numbers = append(numbers, 6)
	numbers = append(numbers, 31)
	numbers = append(numbers, 17)
	numbers = append(numbers, 9)
	numbers = append(numbers, 48)
	numbers = append(numbers, 53)
	assert.Equal(t, numbers, getNumbers("83 86  6 31 17  9 48 53"))
}

func TestGetCard(t *testing.T) {
	card := Card{}
	card.id = 1
	card.winningNumbers = []int{41, 48, 83, 86, 17}
	card.playerNumbers = []int{83, 86, 6, 31, 17, 9, 48, 53}
	assert.Equal(t, card, getCard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"))

	card = Card{}
	card.id = 2
	card.winningNumbers = []int{13, 32, 20, 16, 61}
	card.playerNumbers = []int{61, 30, 68, 82, 17, 32, 24, 19}
	assert.Equal(t, card, getCard("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"))

	card = Card{}
	card.id = 3
	card.winningNumbers = []int{1, 21, 53, 59, 44}
	card.playerNumbers = []int{69, 82, 63, 72, 16, 21, 14, 1}
	assert.Equal(t, card, getCard("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"))

	card = Card{}
	card.id = 4
	card.winningNumbers = []int{41, 92, 73, 84, 69}
	card.playerNumbers = []int{59, 84, 76, 51, 58, 5, 54, 83}
	assert.Equal(t, card, getCard("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"))

	card = Card{}
	card.id = 5
	card.winningNumbers = []int{87, 83, 26, 28, 32}
	card.playerNumbers = []int{88, 30, 70, 12, 93, 22, 82, 36}
	assert.Equal(t, card, getCard("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"))

	card = Card{}
	card.id = 6
	card.winningNumbers = []int{31, 18, 13, 56, 72}
	card.playerNumbers = []int{74, 77, 10, 23, 35, 67, 36, 11}
	assert.Equal(t, card, getCard("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"))
}

func TestGetPlayerWinningCount(t *testing.T) {
	card := Card{}
	card.id = 1
	card.winningNumbers = []int{41, 48, 83, 86, 17}
	card.playerNumbers = []int{83, 86, 6, 31, 17, 9, 48, 53}
	assert.Equal(t, 4, getPlayerWinningCount(card))

	card = Card{}
	card.id = 2
	card.winningNumbers = []int{13, 32, 20, 16, 61}
	card.playerNumbers = []int{61, 30, 68, 82, 17, 32, 24, 19}
	assert.Equal(t, 2, getPlayerWinningCount(card))

	card = Card{}
	card.id = 3
	card.winningNumbers = []int{1, 21, 53, 59, 44}
	card.playerNumbers = []int{69, 82, 63, 72, 16, 21, 14, 1}
	assert.Equal(t, 2, getPlayerWinningCount(card))

	card = Card{}
	card.id = 4
	card.winningNumbers = []int{41, 92, 73, 84, 69}
	card.playerNumbers = []int{59, 84, 76, 51, 58, 5, 54, 83}
	assert.Equal(t, 1, getPlayerWinningCount(card))

	card = Card{}
	card.id = 5
	card.winningNumbers = []int{87, 83, 26, 28, 32}
	card.playerNumbers = []int{88, 30, 70, 12, 93, 22, 82, 36}
	assert.Equal(t, 0, getPlayerWinningCount(card))

	card = Card{}
	card.id = 6
	card.winningNumbers = []int{31, 18, 13, 56, 72}
	card.playerNumbers = []int{74, 77, 10, 23, 35, 67, 36, 11}
	assert.Equal(t, 0, getPlayerWinningCount(card))
}

func TestGetPoints(t *testing.T) {
	assert.Equal(t, 0, getPoints(0))
	assert.Equal(t, 1, getPoints(1))
	assert.Equal(t, 2, getPoints(2))
	assert.Equal(t, 8, getPoints(4))
}
