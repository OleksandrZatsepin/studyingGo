package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var n, s, sum int

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		var err error
		n, err = strconv.Atoi(res)
		if err != nil {
			return
		}
	}
	if n > 0 {
		for i := 2; i < n+2; i++ {
			if input.Scan() {
				res := input.Text()
				var err error
				s, err = strconv.Atoi(res)
				if err != nil {
					return
				}
				if s < 0 {
					return
				}
			}

			if i%2 != 0 {
				sum = sum - s
			} else {
				sum = sum + s
			}
		}
		println(sum)
	}
}

// Почему неверный ответ???
