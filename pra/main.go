package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Iterate through the command-line arguments to ensure flags use "--" and "="
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			fmt.Printf("found: %s\n", arg)
			os.Exit(1)
		}
		if strings.HasPrefix(arg, "--") && !strings.Contains(arg, "=") {
			fmt.Fprintf(os.Stderr, "Error: Flags should use '=' to specify values, found: %s\n", arg)
			os.Exit(1)
		}
	}

	// Define the flag
	nameofp := flag.String("name", "John Doe", "Name of the Person")

	// Parse the flags
	flag.Parse()

	// Print the value of the flag
	fmt.Println("Name:", *nameofp)
}
