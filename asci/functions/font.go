package functions

import (
	"bufio"
	"os"
)

func FontLoader(fontFile string) ([95][8]string, error) {
	fontData := [95][8]string{}

	file, err := os.Open(fontFile)
	if err != nil {
		return [95][8]string{}, err
	}
	scanner := bufio.NewScanner(file)
	character := -1
	line := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			character++
			line = 0
			continue
		}
		fontData[character][line] = scanner.Text()
		line++
	}
	if err := scanner.Err(); err != nil {
		return [95][8]string{}, err
	}
	return fontData, nil
}
