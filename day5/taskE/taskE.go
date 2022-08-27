package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var str []rune

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		str = bytes.Runes([]byte(res))

		if len(str) > 1 {
			for i := 0; i < len(str); i++ {
				if (i)%2 == 0 {
					fmt.Printf("%s%s%s", string(str[i]), string(str[i]), string(str[i]))
				} else {
					continue
				}
			}
		}
	}
}
