package functions

import "fmt"

func AscciArt(r rune, fontData [95][8]string) ([]string, error) {
	line := []string{}
	if r < 32 || r > 126 {
		return nil, fmt.Errorf("ERROR: ")
	}
	char := int(r) - 32
	for i := 0; i < 8; i++ {
		line = append(line, fontData[char][i])
		
	}
	return line, nil
}

