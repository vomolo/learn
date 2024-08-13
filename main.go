package main

import "fmt"

func main() {
	str := "    hello  good     morning   "
	res := TrimSpace1(str)
	res1 := WordSlice(res)
	fmt.Println(res1)
}

func TrimSpace1(str string) string {
	start := 0
	end := len(str) - 1

	for start < end && str[start] == ' ' {
		start++
	}

	for end > start && str[end] == ' ' {
		end--
	}
	return str[start:end]
}

func WordSlice(str string) string {
	var word []rune

	for _, cha := range str {
		if cha != ' ' {
			word = append(word, cha)
		}

		if cha == ' ' {
			fmt.Println()
		}
	}
	return string(word)
}
