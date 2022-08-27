package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var one []rune
	var two []rune

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		one = bytes.Runes([]byte(res))
		if len(one) == 0 {
			return
		}
	}

	for {
		if input.Scan() {
			res := input.Text()
			two = bytes.Runes([]byte(res))
			if len(two) == 0 {
				return
			}
		}
		if two[0] == one[len(one)-1] {
			one = two
			continue
		} else {
			fmt.Print(string(two))
			return
		}
	}
}
