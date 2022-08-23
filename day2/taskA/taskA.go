package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string
	var price = 23

	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}

	n := utf8.RuneCountInString(s)

	price = n * 23

	if price < 100 {
		fmt.Printf("%d коп.", price)
		return
	}

	fmt.Printf("%d р. %d коп.", (price-price%100)/100, price%100)
}
