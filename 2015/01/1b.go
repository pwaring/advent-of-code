package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	floor := 0
	position := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	for _, c := range input {
		str := string(c)

		if str == "(" {
			floor++
		} else if str == ")" {
			floor--
		}

		position++

		if floor < 0 {
			fmt.Println(position)
			os.Exit(0)
		}
	}

	fmt.Println("Santa did not enter the basement at any point")
}
