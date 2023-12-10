package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHand(t *testing.T) {
	hand := Hand{}
	hand.cards = append(hand.cards, Card{
		label: "A",
	})
	hand.cards = append(hand.cards, Card{
		label: "A",
	})
	hand.cards = append(hand.cards, Card{
		label: "A",
	})
	hand.cards = append(hand.cards, Card{
		label: "A",
	})
	hand.cards = append(hand.cards, Card{
		label: "A",
	})
	hand.handType = "FiveKind"
	assert.Equal(t, hand, getHand("AAAAA"))

	hand = Hand{}
	hand.cards = append(hand.cards, Card{
		label: "2",
	})
	hand.cards = append(hand.cards, Card{
		label: "3",
	})
	hand.cards = append(hand.cards, Card{
		label: "4",
	})
	hand.cards = append(hand.cards, Card{
		label: "5",
	})
	hand.cards = append(hand.cards, Card{
		label: "6",
	})
	hand.handType = "HighCard"
	assert.Equal(t, hand, getHand("23456"))
}

func TestGetBid(t *testing.T) {
	bid := Bid{}
	bid.amount = 765
	assert.Equal(t, bid, getBid("765"))

	bid = Bid{}
	bid.amount = 684
	assert.Equal(t, bid, getBid("684"))

	bid = Bid{}
	bid.amount = 28
	assert.Equal(t, bid, getBid("28"))

	bid = Bid{}
	bid.amount = 220
	assert.Equal(t, bid, getBid("220"))

	bid = Bid{}
	bid.amount = 483
	assert.Equal(t, bid, getBid("483"))
}

func TestGetHandType(t *testing.T) {
	assert.Equal(t, "FiveKind", getHandType(getHand("AAAAA")))
	assert.Equal(t, "FourKind", getHandType(getHand("AA8AA")))
	assert.Equal(t, "FullHouse", getHandType(getHand("23332")))
	assert.Equal(t, "ThreeKind", getHandType(getHand("TTT98")))
	assert.Equal(t, "TwoPair", getHandType(getHand("23432")))
	assert.Equal(t, "OnePair", getHandType(getHand("A23A4")))
	assert.Equal(t, "HighCard", getHandType(getHand("23456")))
	assert.Equal(t, "OnePair", getHandType(getHand("32T3K")))
	assert.Equal(t, "TwoPair", getHandType(getHand("KK677")))
	assert.Equal(t, "TwoPair", getHandType(getHand("KTJJT")))
	assert.Equal(t, "ThreeKind", getHandType(getHand("T55J5")))
	assert.Equal(t, "ThreeKind", getHandType(getHand("QQQJA")))
}
