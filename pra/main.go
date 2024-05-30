package main

import "os"

func Wordmatch(s1, s2 string) bool {
	s1 := os.Args[1]
	s2 := os.Args[2]

	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	return i == len(s1)
}
