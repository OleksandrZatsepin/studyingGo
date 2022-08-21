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

	fmt.Printf("Периметр: %d\n", (a+b)*2)
	fmt.Printf("Площадь: %d", a*b)
}
