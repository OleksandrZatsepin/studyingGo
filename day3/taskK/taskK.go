package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var p [][]int
	var s string

	_, err := fmt.Scan("%d", &s)
	if err != nil {
		return
	}
	n, err = strconv.Atoi(s)
	if err != nil {
		return
	}

	m := n
	p[0][0] = 1
	p[1][n-1] = 1
	if n > 1 {
		p[1][m] = n - 1
		p[n-2][] = n - 1
	}

	for i := 2; i < n; i++ {
		p[i] = p[i-2] + p[i-1]

	}
	fmt.Println(p[n-1])
}
