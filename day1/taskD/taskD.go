package main

import "fmt"

func main() {
	var firstName, lastName string
	var age int

	_, err := fmt.Scanln(&firstName)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&lastName)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&age)
	if err != nil {
		return
	}

	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS", firstName, lastName, age)
}
