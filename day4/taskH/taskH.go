package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num []int
	var sum int

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		n, err := strconv.Atoi(res)
		if err != nil || n <= 0 {
			return
		}
		num = append(num, n)
	}
	for i := 0; i < num[0]; i++ {
		if input.Scan() {
			res := input.Text()
			n, err := strconv.Atoi(res)
			if err != nil {
				return
			}
			num = append(num, n)
		}
	}
	if input.Scan() {
		res := input.Text()
		a, err := strconv.Atoi(res)
		if err != nil || a <= 0 {
			return
		}
		if a <= num[0] {
			num = append(num, a)
		} else {
			return
		}
	}

	if input.Scan() {
		res := input.Text()
		b, err := strconv.Atoi(res)
		if err != nil || b <= 0 {
			return
		}
		if b <= num[0] && b >= num[len(num)-1] {
			num = append(num, b)
		} else {
			return
		}
	}

	for i := num[len(num)-2]; i <= num[len(num)-1]; i++ {
		sum += num[i]
	}
	fmt.Print(sum)
}
