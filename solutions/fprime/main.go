package main

import (
	"fmt"
	"os"
	"strconv"
)

// func fprime(value int) {
// 	if value == 1 {
// 		return
// 	}
// 	div := 2
// 	for value > 1 {
// 		if value%div == 0 {
// 			fmt.Print(div)
// 			value = value / div

// 			if value > 1 {
// 				fmt.Print("*")
// 			}
// 			div--
// 		}
// 		div++
// 	}
// 	fmt.Println()
// }

func main() {
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			fprime(i)
		}
	}
}

func fprime(val int) {
	if val == 1 {
		return
	}

	div := 2
	for val > 1 {
		if val%div == 0 {
			fmt.Print(div)
			val = val / div
			if val > 1 {
				fmt.Print("*")
			}
			div--
		}
		div++
	}
	fmt.Println()
}
