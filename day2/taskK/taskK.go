package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}

	if strings.Contains(s, "чат") {
		fmt.Print("БОТ")
		return
	}
	fmt.Print("НЕ БОТ")
}
