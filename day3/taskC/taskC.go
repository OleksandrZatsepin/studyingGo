package main

import "fmt"

func main() {
	var n uint

	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			return
		}

		if n == 0 {
			return
		}
		fmt.Println(n)
	}
}
