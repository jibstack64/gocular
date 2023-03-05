package gocular

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Clears the previous line from the console.
func ClearLine() {
	fmt.Printf("\x1b[1A\x1b[2K")
}

// Clears `x` lines from the console.
func ClearLines(x int) {
	for i := 0; i < x; i++ {
		ClearLine()
	}
}

// Clears the entire console.
func Clear() {
	if runtime.GOOS == "windows" {
		exec.Command("cls").Run()
	} else {
		exec.Command("clear").Run()
	}
}
