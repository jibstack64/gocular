package gocular

import (
	"fmt"
)

// Clears the previous line from the console.
func ClearLine() {
	fmt.Printf("\x1b[1A\x1b[2K")
}
