package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNode(t *testing.T) {
	assert.Equal(t, Node{
		label: "AAA",
		left:  "BBB",
		right: "CCC",
	}, getNode("AAA = (BBB, CCC)"))

	assert.Equal(t, Node{
		label: "BBB",
		left:  "DDD",
		right: "EEE",
	}, getNode("BBB = (DDD, EEE)"))

	assert.Equal(t, Node{
		label: "CCC",
		left:  "ZZZ",
		right: "GGG",
	}, getNode("CCC = (ZZZ, GGG)"))

	assert.Equal(t, Node{
		label: "DDD",
		left:  "DDD",
		right: "DDD",
	}, getNode("DDD = (DDD, DDD)"))

	assert.Equal(t, Node{
		label: "EEE",
		left:  "EEE",
		right: "EEE",
	}, getNode("EEE = (EEE, EEE)"))

	assert.Equal(t, Node{
		label: "GGG",
		left:  "GGG",
		right: "GGG",
	}, getNode("GGG = (GGG, GGG)"))

	assert.Equal(t, Node{
		label: "ZZZ",
		left:  "ZZZ",
		right: "ZZZ",
	}, getNode("ZZZ = (ZZZ, ZZZ)"))
}

func TestFindNode(t *testing.T) {
	nodes := []Node{}
	nodes = append(nodes, getNode("AAA = (BBB, CCC)"))
	nodes = append(nodes, getNode("BBB = (DDD, EEE)"))
	nodes = append(nodes, getNode("CCC = (ZZZ, GGG)"))
	nodes = append(nodes, getNode("DDD = (DDD, DDD)"))
	nodes = append(nodes, getNode("EEE = (EEE, EEE)"))
	nodes = append(nodes, getNode("GGG = (GGG, GGG)"))
	nodes = append(nodes, getNode("ZZZ = (ZZZ, ZZZ)"))

	assert.Equal(t, &nodes[0], findNode("AAA", nodes))
	assert.Equal(t, &nodes[1], findNode("BBB", nodes))
	assert.Equal(t, &nodes[2], findNode("CCC", nodes))
	assert.Equal(t, &nodes[3], findNode("DDD", nodes))
	assert.Equal(t, &nodes[4], findNode("EEE", nodes))
	assert.Equal(t, &nodes[5], findNode("GGG", nodes))
	assert.Equal(t, &nodes[6], findNode("ZZZ", nodes))
	assert.Nil(t, findNode("AAB", nodes))
	assert.Nil(t, findNode("FFF", nodes))
	assert.Nil(t, findNode("123", nodes))
}
