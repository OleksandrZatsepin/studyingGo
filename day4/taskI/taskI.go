package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var job []string
	var num []int

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
			job = append(job, res)
		}
	}

	if input.Scan() {
		res := input.Text()
		m, err := strconv.Atoi(res)
		if err != nil || m <= 0 {
			return
		}
		num = append(num, m)
	}

	for i := 0; i < num[1]; i++ {
		if input.Scan() {
			res := input.Text()
			n, err := strconv.Atoi(res)
			if err != nil || n <= 0 {
				return
			}
			num = append(num, n)
		}
	}

	for i := 0; i < num[1]; i++ {
		fmt.Println(job[num[i+2]-1])
	}
}
