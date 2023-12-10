package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var handTypes = map[string]int{
	"HighCard":  1,
	"OnePair":   2,
	"TwoPair":   3,
	"ThreeKind": 4,
	"FullHouse": 5,
	"FourKind":  6,
	"FiveKind":  7,
}

type Card struct {
	label string
}

type Hand struct {
	cards    []Card
	handType string
}

type Bid struct {
	amount int
}

type Round struct {
	hand Hand
	bid  Bid
}

func getRound(str string) Round {
	round := Round{}
	str = strings.TrimSpace(str)
	parts := strings.Fields(str)

	round.hand = getHand(parts[0])
	round.bid = getBid(parts[1])

	return round
}

func getBid(str string) Bid {
	bid := Bid{}
	str = strings.TrimSpace(str)
	bid.amount, _ = strconv.Atoi(str)

	return bid
}

func getHandType(hand Hand) string {
	// Default is HighCard as all hands must have a type
	handType := "HighCard"

	cardCounts := map[Card]int{}

	for c := range hand.cards {
		cardCounts[hand.cards[c]]++
	}

	uniqueCards := []int{}
	for cc := range cardCounts {
		uniqueCards = append(uniqueCards, cardCounts[cc])
	}

	if len(uniqueCards) == 1 {
		// Only one uniqe card, so they must all be the same
		handType = "FiveKind"
	} else if len(uniqueCards) == 2 {
		// Must be Four of a Kind or Full House if we have two unique cards
		if slices.Contains(uniqueCards, 4) {
			handType = "FourKind"
		} else {
			handType = "FullHouse"
		}
	} else if len(uniqueCards) == 3 {
		// Must be Three of a Kind or Two Pair if we have three unique cards
		if slices.Contains(uniqueCards, 3) {
			handType = "ThreeKind"
		} else {
			handType = "TwoPair"
		}
	} else if len(uniqueCards) == 4 {
		handType = "OnePair"
	}

	return handType
}

func getHand(str string) Hand {
	hand := Hand{}

	str = strings.TrimSpace(str)
	cardLabels := strings.Split(str, "")

	for cl := range cardLabels {
		hand.cards = append(hand.cards, Card{
			label: cardLabels[cl],
		})
	}

	hand.handType = getHandType(hand)

	return hand
}

// Sort rounds from weakest to strongest
func sortRounds(rounds []Round) []Round {
	sort.Slice(rounds, func(x, y int) bool {
		if handTypes[rounds[x].hand.handType] < handTypes[rounds[y].hand.handType] {
			return true
		} else if handTypes[rounds[x].hand.handType] > handTypes[rounds[y].hand.handType] {
			return false
		} else {
			// Hands have the same type, so we need to compare each card in order
			for c := range rounds[x].hand.cards {
				if cardValues[rounds[x].hand.cards[c].label] < cardValues[rounds[y].hand.cards[c].label] {
					return true
				} else if cardValues[rounds[x].hand.cards[c].label] > cardValues[rounds[y].hand.cards[c].label] {
					return false
				}
			}

			// If we get this far, both hands are equal and the other doesn't
			// matter, but we have to return something
			return false
		}
	})

	return rounds
}

func main() {
	rounds := []Round{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			rounds = append(rounds, getRound(line))
		}
	}

	// Sort rounds by score
	rounds = sortRounds(rounds)

	winnings := 0
	for r := range rounds {
		// Slices are 0-indexed but ranks are 1-indexed
		rank := r + 1
		winnings += rank * rounds[r].bid.amount
	}

	fmt.Println(winnings)
}
