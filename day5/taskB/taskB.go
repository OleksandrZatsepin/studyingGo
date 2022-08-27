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

		if (str[0] == 1044 || str[0] == 1076) && (str[len(str)-1] == 1040 || str[len(str)-1] == 1072) || (str[0] == 1040 || str[0] == 1072) && (str[len(str)-1] == 1044 || str[len(str)-1] == 1076) {
			fmt.Println("СОГЛАСЕН")
		} else {
			fmt.Println("НЕ СОГЛАСЕН")
		}
	}
}
