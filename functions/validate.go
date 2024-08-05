package functions

import (
	"fmt"
	"unicode"
)

// Checks if a string contains non-ASCII characters i.e characters >127 in the ascii table.
func ContainsNonASCII(s string) error {
	for _, char := range s {
		if char > unicode.MaxASCII {
			err := fmt.Errorf("400: non ascii character")
			return err
		}
	}
	return nil
}
