package main

import "fmt"

func main() {
	var a, b uint

	_, err := fmt.Scanln(&a)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&b)
	if err != nil {
		return
	}

	fmt.Print((a + b) * (a + b))
}
