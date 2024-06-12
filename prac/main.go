package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . --color=red kit 'a king kitten have kit'")
        return
    }
    color := os.Args[1]
    text := os.Args[2]
    if color == "--color=red" {
        for _, word := range strings.Split(text, " ") {
            if word == "kitten" || word == "kit" {
                fmt.Printf("\033[91m%s\033[0m\n", word)
            } else {
                fmt.Println(word)
            }
        }
    } else {
        fmt.Println("Invalid color. Supported colors are red.")
    }
}

