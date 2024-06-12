package functions

import (
	"fmt"
	//"strings"
)

func SubstringColoring(lines []string, color string) []string {
	// colorCode := fmt.Sprintf("\033[%sm", color)
	// resetCode := "\033[0m"

	for i := range lines{
		lines[i]=fmt.Sprintf("\033[%sm%s33[0m",color,lines[i])
	}
	return lines
}
