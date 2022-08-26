package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string

	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}

	n, err1 := strconv.Atoi(s)
	if err1 != nil {
		return
	}

	if n > 0 {
		fmt.Printf("%b", n)
	}
}
