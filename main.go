package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	//args1 := os.Args[1]
	//args2 := os.Args[2]

	res1 := sliceString(args)
	res2 := lowerCaps(res1)
	res3 := lastCap(res2)
	// args3 := os.Args[3]
	fmt.Println(args)
	//fmt.Println(args1)
	//fmt.Println(args2)
	fmt.Println(res3)
	// fmt.Println(args3)
}

func sliceString(str []string) string {
	res := ""

	for _, char := range str {
		res += string(char) + "\n"
	}
	return res

}

func lowerCaps(str string) string {
	res := ""

	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
			res += string(char)
		} else {

			res += string(char)
		}

	}
	return res
}

func lastCap(str string) string {
	res := ""
	n := len(str)

	for i := 0; i < n; i++ {
		char := str[i]
		// Capitalize the character if it's the last character of a word and is lowercase
		if (i == n-1 || str[i+1] == ' ') && char >= 'a' && char <= 'z' {
			char = char - 32
		}
		res += string(char) // Append the character to the result
	}
	return res
}
