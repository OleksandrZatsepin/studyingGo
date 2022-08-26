package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var a []int

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		res := input.Text()
		n, err := strconv.Atoi(res)
		if err != nil {
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

			fmt.Printf("%d", a[0])
			for i := 1; i < len(a); i++ {
				fmt.Printf(" %d", a[i])
			}

			if len(a) == 2 {
				fmt.Print("\nACHTUNG")
			}
		}
	}
}
