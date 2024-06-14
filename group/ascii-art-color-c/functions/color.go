package functions

import (
	"fmt"
	"log"
)

func ColoredSub(substr string, fontarray [95][8]string, colorcode string) []string {
	concatenatedLines := make([]string, 8)

	for _, char := range substr {
		Ascii, err := PrintChar(char, fontarray)
		if err != nil {
			log.Fatal(err)
		}
		for i, line := range Ascii {
			if colorcode != "" {
				concatenatedLines[i] += fmt.Sprintf("\033[%sm%s\033[0m", colorcode, line)
			} else {
				concatenatedLines[i] += line
			}
		}
	}
	return concatenatedLines
}
