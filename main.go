// package main

// import "fmt"

// func main() {
// 	str := "    hello  good     morning   "
// 	res := TrimSpace1(str)
// 	res1 := WordSlice(res)
// 	fmt.Println(res1)
// }

// func TrimSpace1(str string) string {
// 	start := 0
// 	end := len(str) - 1

// 	for start < end && str[start] == ' ' {
// 		start++
// 	}

// 	for end > start && str[end] == ' ' {
// 		end--
// 	}
// 	return str[start:end]
// }

// func WordSlice(str string) string {
// 	var word []rune

// 	for _, cha := range str {
// 		if cha != ' ' {
// 			word = append(word, cha)
// 		}

// 		if cha == ' ' {
// 			fmt.Println()
// 		}
// 	}
// 	return string(word)
// }

package main

import "fmt"

func WordSlice(str string) string {
	var word []rune
	var result []rune

	for i, char := range str {
		if char != ' ' {
			word = append(word, char)
		} else if len(word) > 0 {
			result = append(result, word...)
			result = append(result, '   ')
			word = nil
		}

		if i == len(str)-1 && len(word) > 0 {
			result = append(result, word...)
		}
	}

	return string(result)
}

func main() {
	input := "   Hello       World!     This is     a    test.   "
	result := WordSlice(input)
	fmt.Println(result) // Output: "Hello  World!  This  is  a  test."
}
