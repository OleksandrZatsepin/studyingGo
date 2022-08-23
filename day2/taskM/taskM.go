package main

import (
	"fmt"
)

func main() {
	var s1, s2, s3 string

	_, err := fmt.Scanln(&s1)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&s2)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&s3)
	if err != nil {
		return
	}

	if s1 == "1" && s2 == "2" && s3 == "3" {
		fmt.Print("ОК")
		return
	}

	if (s1 == "раз" || s1 == "один") && s2 == "два" && s3 == "три" {
		fmt.Print("ОК")
		return
	}

	fmt.Print("НЕ ПРАВИЛЬНО")
}
