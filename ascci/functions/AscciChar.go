package functions

func AsciiChar(C rune, fontData[95][8]string)[]string{
	var lines []string
	//characters in Ascii begin from index 32 so we subtract to get the right indexing
	char:=int(C)-32

	//loop through each line in the character and add it to the lines array

	for i:=0;i<=7;i++{
		lines = append(lines, fontData[char][i])
	}

	return lines
}