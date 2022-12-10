package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const chunkLength = 4

	// ASSUMPTION: There is only one line of input and no errors
	byteInput, _ := os.ReadFile(os.Stdin.Name())
	input := string(byteInput)

	for start, end := 0, chunkLength-1; end < len(input); start, end = start+1, end+1 {
		// Add one to end because second array slice index is exclusive
		chunk := input[start : end+1]
		characters := strings.Split(chunk, "")
		seen := make(map[string]bool)

		for c := range characters {
			seen[characters[c]] = true
		}

		if len(seen) == chunkLength {
			// All characters are unique
			fmt.Println(end + 1)
			os.Exit(0)
		}
	}
}
