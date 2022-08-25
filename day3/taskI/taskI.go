package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s string
	var a []int

	_, err := fmt.Scan(&s)
	if err != nil {
		return
	}

	n, err1 := strconv.Atoi(s)
	if err1 != nil {
		return
	}

	if n > 0 && n < 1000000 {
		for i := 1; i*i <= n; i++ {
			if n%i == 0 {
				a = append(a, i)
				if i != n/i {
					a = append(a, n/i)
				}
			}
		}
		sort.Ints(a)
		for i := 0; i < len(a); i++ {
			fmt.Printf("%d ", a[i])
		}

		if len(a) == 2 {
			print("\nACHTUNG")
		}
	}
}

// Почему неверный ответ???
