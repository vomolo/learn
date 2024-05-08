package tez

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Hex converts hexadecimal numbers in a string to their decimal equivalents.
func Hex(line string) string {
	hex := regexp.MustCompile(`(?i)(\b\w+\b)\s?\(hex\)`)
	noSpaceHex := regexp.MustCompile(`(?i)(\b\w+\b)\s?[\(]hex[\)]`)
	for {
		matchHex := hex.FindStringSubmatch(line)
		matchNoSpaceHex := noSpaceHex.FindStringSubmatch(line)
		if len(matchHex) == 0 && len(matchNoSpaceHex) == 0 {
			break
		}
		var hexNumber string
		if len(matchHex) > 0 {
			hexNumber = matchHex[1]
		} else if len(matchNoSpaceHex) > 0 {
			hexNumber = matchNoSpaceHex[1]
		}
		decimalHex, err := strconv.ParseInt(hexNumber, 16, 64)
		if err != nil {
			// Report the error but continue processing
			// fmt.Printf("Error converting %s to decimal: %v\n", hexNumber, err)
			line = strings.Replace(line, fmt.Sprintf("%s (hex)", hexNumber), "ERROR", 1)
			line = strings.Replace(line, fmt.Sprintf("%s(hex)", hexNumber), "ERROR", 1)
			fields := strings.Fields(line)
			line = strings.Join(fields, " ")
			continue
		}
		line = strings.Replace(line, fmt.Sprintf("%s (hex)", hexNumber), fmt.Sprintf("%d", decimalHex), 1)
		line = strings.Replace(line, fmt.Sprintf("%s(hex)", hexNumber), fmt.Sprintf("%d", decimalHex), 1)
		fields := strings.Fields(line)
		line = strings.Join(fields, " ")
	}
	return line
}
