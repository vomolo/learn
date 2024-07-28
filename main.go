package main

import (
	"fmt"
)

func Atoi(s string) (int, error) {
	if len(s) < 0 {
		return 0, fmt.Errorf("String is Empty")
	}
}
