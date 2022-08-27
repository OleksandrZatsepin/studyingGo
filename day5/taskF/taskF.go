package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str []rune
	var start, stop int

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()

		if strings.ContainsRune(res, 9) || strings.ContainsRune(res, 32) {
			return
		}

		str = bytes.Runes([]byte(res))

		start = 0
		stop = len(str)

		if len(str) > 2 {
			for stop-start > 0 {
				for j := start; j < stop; j++ {
					fmt.Print(string(str[j]))
				}
				start = start + 2

				if stop-start > 0 {
					fmt.Print("\n")
				}

				for j := start; j < stop; j++ {
					fmt.Print(string(str[j]))
				}

				stop = stop - 2

				if stop-start > 0 {
					fmt.Print("\n")
				}
			}

		}
	}
}

//  Test #8 - not passed.
