package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	seen := make(map[rune]bool)
	store := []rune{}

	for _, char1 := range str1 {
		if !seen[char1] {
			store = append(store, char1)
			seen[char1] = true
		}
	}

	for _, char2 := range str2 {
		if !seen[char2] {
			store = append(store, char2)
			seen[char2] = true
		}
	}
	res := string(store)
	return len(res)
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
