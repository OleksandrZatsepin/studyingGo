package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 uint

	_, err := fmt.Scanln(&x1)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&y1)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&x2)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&y2)
	if err != nil {
		return
	}

	if x1 > 8 || y1 > 8 || x2 > 8 || y2 > 8 {
		return
	}

	if (x2 == x1-2 || x2 == x1+2) && (y2 == y1-1 || y2 == y1+1) || (x2 == x1-1 || x2 == x1+1) && (y2 == y1-2 || y2 == y1+2) {
		fmt.Print("ДА")
		return
	}

	fmt.Print("НЕТ")
}
