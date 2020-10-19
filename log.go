package stat

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

// Success prints a successful message to the console
// and prepends it with [SCSS] in bold green text.
func Success(message string) {
	fmt.Println(aurora.Green(aurora.Bold("[SCSS]")), message)
}

// Info prints an informational message to the console
// and prepends it with [INFO] in bold cyan text.
func Info(message string) {
	fmt.Println(aurora.Cyan(aurora.Bold("[INFO]")), message)
}

// Warn prints a warning message to the console
// and prepends it with [WARN] in bold yellow text.
func Warn(message string) {
	fmt.Println(aurora.Yellow(aurora.Bold("[WARN]")), message)
}

// Fail prints an error message to the console
// and prepends it with [FAIL] in bold red text.
// It also exits the program with an exit code of 1.
func Fail(message string) {
	fmt.Println(aurora.Red(aurora.Bold("[FATL]")), message)
	os.Exit(1)
}
