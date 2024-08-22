package main

import "fmt"

func TrimSpaces(str string) string {
	start := 0
	end := len(str) - 1

	for start <= end && str[start] == ' ' {
		start++
	}

	for end >= start && str[end] == ' ' {
		end--
	}

	return str[start : end+1]
}

func RemSpaces(str string) string {
	res := ""

	for _, char := range str {
		if char != ' ' {
			res += string(char)
		} else if char == ' ' {
			res += " "
		}
	}
	return res
}

func UnderScore(str string) string {
	res := ""

	for _, char := range str {
		if char != ' ' {
			res += string(char)
		} else if char == ' ' {
			res += "_"
		}
	}
	return res
}

func UnderscoreCapital(str string) string {
	res := ""

	for i, char := range str {
		if char >= 'A' && char <= 'Z' && i > 0 && str[i-1] != '_' {
			res += "_" // Add underscore before the capital letter
		}
		res += string(char)
	}

	return res
}

func main() {
	str := "   hey JudeTimeYYes  "
	trim := TrimSpaces(str)
	rem := RemSpaces(trim)
	fmt.Println(rem)
	under := UnderScore(rem)
	un := UnderscoreCapital(under)
	fmt.Println(un)
}
