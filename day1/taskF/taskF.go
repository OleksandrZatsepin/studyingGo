package main

import "fmt"

func main() {
	var word [4]string
	var author string

	for i := 0; i < 4; i++ {
		_, err := fmt.Scanln(&word[i])
		if err != nil {
			return
		}
	}
	_, err := fmt.Scanln(&author)
	if err != nil {
		return
	}

	for i := 3; i >= 0; i-- {
		fmt.Printf("%s - %s\n", word[i], author)
	}
}
