package main

import "fmt"

func Atoi(s string) int {
	var res int
	sign := 1
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}

	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		sign = 1
		i++
	}

	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return 0
		} else {
			res = res*10 + int(s[i]-'0')
			i++
		}
	}
	return res * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
