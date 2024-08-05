package functions

import (
	"strings"
)

/*
This function is responsible for converting a given input text into ASCII art after checking for special characters in the input . The function  calls the other function `PrintLineByLine`, to do the actual ASCII printing.
*/
func PrintAsciiArt(text string, asciiArt []string) string {
	var builder1 strings.Builder
	if text == "" {
		return ""
	}

	// handling cases of string literals
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			_, err := builder1.WriteString(PrintLineByLine(line, asciiArt))
			if err != nil {
				return ""
			}
		}
	}
	return builder1.String()
}
