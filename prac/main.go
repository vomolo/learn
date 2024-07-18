package main

import (
	"fmt"
)

func main() {
	var str string
	for i := 'a'; i <= 'z'; i++ {
		str += string(i)
	}
	fmt.Println(str)
}
