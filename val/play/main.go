package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("enter name")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("my name is %v, you all\n", name)
}
