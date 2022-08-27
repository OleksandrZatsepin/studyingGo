package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var str []rune
	var count, max int

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		str = bytes.Runes([]byte(res))
	}

	if len(str) == 0 {
		return
	}

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 1086:
			count = count + 1
			if max < count {
				max = count
			}
		case 1088:
			if max < count {
				max = count
			}
			count = 0
		default:
			return
		}
	}
	fmt.Print(max)
}
