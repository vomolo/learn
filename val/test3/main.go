package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Up(str string) string {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(up)" {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return strings.Join(slice, " ")
}

func UpInt(str string) (string, error) {
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if slice[i] == "(up," {
			if i+1 >= len(slice) {
				return "", fmt.Errorf("index out of range: %d", i+1)
			}
			num, err := strconv.Atoi(slice[i+1][:len(slice[i+1])-1])
			if err != nil {
				return "", err
			}
			if i-num < 0 {
				return "", fmt.Errorf("index out of range: %d", i-num)
			}
			for j := i - num; j < i; j++ {
				slice[j] = strings.ToUpper(slice[j])
			}
			slice = append(slice[:i], slice[i+2:]...)
		}
	}
	return strings.Join(slice, " "), nil
}

func main() {
	sen := "If I talk there will be (up, 7) fire in hell"
	res := Up(sen)
	result, err := UpInt(res)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
