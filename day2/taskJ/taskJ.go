package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		return
	}

	//Wrong Answer
	r := a + b
	_, frac := math.Modf(r)

	if frac == 0 {
		if int(r)%2 == 0 {
			fmt.Print("ЧЁТНОЕ")
			return
		}
		fmt.Print("НЕЧЁТНОЕ")
	}
}

//1.5000000000000000000000000000000000000000000000000001
//2.5
//ЧЁТНОЕ
