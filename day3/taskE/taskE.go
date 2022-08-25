package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var pass, pass1 string

	for {
		_, err := fmt.Scanln(&pass)
		if err != nil {
			return
		}
		_, err = fmt.Scanln(&pass1)
		if err != nil {
			return
		}

		if utf8.RuneCountInString(pass) < 8 {
			fmt.Println("Слишком короткий пароль!")
			continue
		}
		if strings.Contains(pass, "123") || strings.Contains(pass, "qwe") {
			fmt.Println("Слишком простой пароль!")
			continue
		}
		if pass != pass1 {
			fmt.Println("Введенные пароли различаются!")
			continue
		}
		fmt.Println("Ну наконец-то!")
		return
	}
}
