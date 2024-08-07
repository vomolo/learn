package solutions

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[0:]

	if len(arg) == 2 {

		args := os.Args[1]
		las := len(args) - 1

		var lastWord string

		for i := las; i >= 0; i-- {
			if args[i] == ' ' && lastWord == "" {
				continue
			}
			if args[i] == ' ' && lastWord != "" {
				break
			}
			{
				lastWord = string(args[i]) + lastWord
			}
		}
		if lastWord == "" {
			return
		}
		fmt.Println(lastWord)
	}
}
