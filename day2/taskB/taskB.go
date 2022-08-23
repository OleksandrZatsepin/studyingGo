package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	_, err := fmt.Scanln(&a)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&b)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&c)
	if err != nil {
		return
	}

	if a == 0 && b == 0 {
		fmt.Print("корней нет")
		return
	}

	if a == 0 && b != 0 {
		fmt.Print("один корень")
		return
	}

	if a != 0 {
		d := b*b - 4*a*c
		if d < 0 {
			fmt.Print("корней нет")
			return
		}
		if d == 0 {
			fmt.Print("один корень")
			return
		}
		if d > 0 {
			fmt.Print("два корня")
			return
		}
	}
}
