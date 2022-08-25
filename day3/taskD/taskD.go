package main

import "fmt"

func main() {
	var s string

	for {
		_, err := fmt.Scanln(&s)
		if err != nil {
			return
		}

		if s == "" {
			return
		}
		fmt.Println(s)
	}
}
