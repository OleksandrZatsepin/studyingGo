package main

import "fmt"

func main() {
	var a, b, c uint

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

	fmt.Printf("%d%d%d", c, b, a)
}
