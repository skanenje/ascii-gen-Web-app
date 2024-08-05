package functions

import (
	"fmt"
	"strings"
)

// extracts the index (ascii index) of the individual characters of the command line argument to match it with the positional index of that letter in the slice which we had stored the ascii art from the ascii banner file
func PrintLineByLine(text string, asciiArt []string) string {
	var builder strings.Builder
	for i := 0; i < asciiArtHeight; i++ {
		for _, char := range text {
			// Convert special characters to their corresponding index
			index := strings.IndexAny(" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", string(char))
			if index != -1 && index < len(asciiArt) {
				_, err := builder.WriteString(strings.Split(asciiArt[index], "\n")[i]) // Print the ith line of the ASCII art for the current character
				if err != nil {
					fmt.Printf("error is %v", err)
				}
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
