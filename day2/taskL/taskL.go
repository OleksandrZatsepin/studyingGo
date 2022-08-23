package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var login, email string

	_, err := fmt.Scanln(&login)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&email)
	if err != nil {
		return
	}

	if strings.Contains(login, "@") || utf8.RuneCountInString(login) < 10 {
		fmt.Print("Некорректный логин")
		return
	}

	//Work in go1.18 and more
	before, after, found := strings.Cut(email, "@")
	if !found || strings.Contains(after, "@") || utf8.RuneCountInString(before) == 0 || strings.Count(after, ".") != 1 {
		fmt.Print("Некорректный email")
		return
	}

	fmt.Print("ОК")
}
