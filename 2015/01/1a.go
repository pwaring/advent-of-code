package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	floor := 0
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
	}

	fmt.Println(floor)
}
