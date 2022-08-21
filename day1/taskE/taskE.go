package main

import "fmt"

func main() {
	var drink, meal, subMeal, time string

	_, err := fmt.Scanln(&drink)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&meal)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&subMeal)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&time)
	if err != nil {
		return
	}

	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'", drink, meal, subMeal, time)
}
