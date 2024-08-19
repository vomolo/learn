// package main

// import "fmt"

// func main() {
// 	// Original string
// 	sentence := "  Go is    an open-source programming language   "

// 	// Splitting the string into a slice of words
// 	words := splitString(sentence)

// 	// Ranging over the slice of words
// 	// for index, word := range words {
// 	// 	fmt.Printf("Word %d: %s\n", index, word)
// 	// }

// 	res := joinWithSpaces(words)
// 	fmt.Println(res)
// }

// func splitString(str string) []string {

// 	var words []string
// 	var currentWord []rune

// 	for _, char := range str {
// 		if char == ' ' {
// 			if len(currentWord) > 0 {
// 				words = append(words, string(currentWord))
// 				currentWord = currentWord[:0]
// 			}
// 		} else {
// 			currentWord = append(currentWord, char)
// 		}
// 	}
// 	if len(currentWord) > 0 {
// 		words = append(words, string(currentWord))
// 	}

// 	return words
// }

// func joinWithSpaces(words []string) string {
// 	var result string

// 	// Join the words with three spaces
// 	for i, word := range words {
// 		if i > 0 {
// 			result += "   "
// 		}
// 		result += word
// 	}

// 	return result
// }

package main

import "fmt"

func TrimSpaces(str string) string {
	start := 0
	end := len(str) - 1

	for start < end && str[start] == ' ' {
		start++
	}

	for end > start && str[end] == ' ' {
		end--
	}

	return str[start : end+1]
}

func InputSpaces(str string) string {
	res := ""
	spaceFlag := false

	for _, char := range str {
		if char != ' ' {
			res += string(char)
			spaceFlag = false
		}else if !spaceFlag {
			res += " "
			spaceFlag =true
		}
	}
	return res
}

func main() {
	str := "    hello       you       there.      "
	res1 := TrimSpaces(str)
	res := InputSpaces(res1)
	fmt.Println(res)
}


package main

import (
    "fmt"
)

// Contains checks if substr is within str.
func Contains(str, substr string) bool {
    // If substr is an empty string, return true
    if substr == "" {
        return true
    }
    
    // Lengths of the strings
    strLen := len(str)
    substrLen := len(substr)

    // If the substring is longer than the string, it cannot be contained
    if substrLen > strLen {
        return false
    }

    // Iterate through the string
    for i := 0; i <= strLen-substrLen; i++ {
        // Check if the substring matches
        if str[i:i+substrLen] == substr {
            return true
        }
    }

    return false
}

func main() {
    fmt.Println(Contains("hello world", "world")) // true
    fmt.Println(Contains("hello world", "golang")) // false
    fmt.Println(Contains("hello world", ""))       // true (empty substring)
}