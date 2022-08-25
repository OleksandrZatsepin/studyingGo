package main

import (
	"fmt"
)

func main() {
	var pulse, min, max float64
	var n uint

	min = 140
	max = 100

	for {
		_, err := fmt.Scanln(&pulse)
		if err != nil {
			return
		}

		if pulse < 0 {
			if n > 0 {
				fmt.Printf("%d\n", n)
				fmt.Printf("%.1f %.1f", min, max)
			}
			return
		}

		if pulse < 100 || pulse > 140 {
			continue
		}

		n += 1
		if pulse < min {
			min = pulse
		}
		if pulse > max {
			max = pulse
		}
	}
}
