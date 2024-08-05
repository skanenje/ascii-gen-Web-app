package functions

func FunctionMain(text, filename string) string {
	asciiArt := ReadAsciiArt(filename)
	s := PrintAsciiArt(text, asciiArt)
	return s
}
