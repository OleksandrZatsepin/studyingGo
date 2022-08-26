package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	var list []string
	var query []string
	var flag bool

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		var err error
		n, err = strconv.Atoi(res)
		if err != nil {
			return
		}
	}
	if input.Scan() {
		res := input.Text()
		var err error
		m, err = strconv.Atoi(res)
		if err != nil {
			return
		}
	}

	for i := 0; i < n; i++ {
		if input.Scan() {
			res := input.Text()
			list = append(list, res)
		}
	}

	for i := 0; i < m; i++ {
		if input.Scan() {
			res := input.Text()
			query = append(query, res)
		}
	}

	for i := 0; i < m; i++ {
		flag = false
		for j := 0; j < n; j++ {
			if strings.Compare(list[j], query[i]) == 0 {
				flag = true
				break
			}
		}

		if flag {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}
}
