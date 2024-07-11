package main

func rot14(c rune) rune {
	if c >= 'A' && c < 'M' || c >= 'a' && c < 'm' {
		return c + 14
	}
	if c >= 'M' && c <= 'Z' || c >= 'm' && c <= 'z' {
		return c - 12
	}
	return c
}

func Rot14(str string) string {
	result := ""
	for _, re := range str {
		result += string(rot14(re))
	}
	return result
}
