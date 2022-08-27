package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var menu []string
	var max int

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		var err error
		max, err = strconv.Atoi(res)
		if err != nil || max == 0 {
			return
		}
	}

	for i := 0; i < max; i++ {
		if input.Scan() {
			res := input.Text()
			//food := bytes.Runes([]byte(res))
			menu = append(menu, res)
		}
	}

}

// Not finished yet.
