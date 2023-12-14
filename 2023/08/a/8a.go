package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	label string
	left  string
	right string
}

func getNode(str string) Node {
	node := Node{}

	str = strings.TrimSpace(str)

	re := regexp.MustCompile(`^(?P<label>[A-Z]+)\s+\=\s+\((?P<left>[A-Z]+),\s+(?P<right>[A-Z]+)\)$`)
	matches := re.FindStringSubmatch(str)
	node.label = matches[re.SubexpIndex("label")]
	node.left = matches[re.SubexpIndex("left")]
	node.right = matches[re.SubexpIndex("right")]

	return node
}

func findNode(index string, nodes []Node) *Node {
	for n := range nodes {
		if nodes[n].label == index {
			return &nodes[n]
		}
	}

	return nil
}

func main() {
	nodes := []Node{}
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)

	parts := strings.Split(inputString, "\n\n")

	// First part is the instructions
	instructions := strings.Split(strings.TrimSpace(parts[0]), "")
	nodeLines := strings.Split(parts[1], "\n")

	for nl := range nodeLines {
		if len(nodeLines[nl]) > 0 {
			nodes = append(nodes, getNode(nodeLines[nl]))
		}
	}

	stepCount := 0
	instructionIndex := 0

	for currentNode := findNode("AAA", nodes); currentNode != nil && currentNode.label != "ZZZ"; stepCount++ {
		if instructionIndex == len(instructions) {
			instructionIndex = 0
		}

		if instructions[instructionIndex] == "L" {
			currentNode = findNode(currentNode.left, nodes)
		} else {
			currentNode = findNode(currentNode.right, nodes)
		}

		instructionIndex++
	}

	fmt.Println(stepCount)
}
