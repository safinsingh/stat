package stat

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

// Success prints a successful message to the console
// and prepends it with [SCSS] in bold green text.
func Success(message ...interface{}) {
	fmt.Println(
		aurora.Green(
			aurora.Bold("[SCSS]"),
		),
		fmt.Sprint(message...),
	)
}

// Info prints an informational message to the console
// and prepends it with [INFO] in bold cyan text.
func Info(message ...interface{}) {
	fmt.Println(
		aurora.Cyan(
			aurora.Bold("[INFO]"),
		),
		fmt.Sprint(message...),
	)
}

// Warn prints a warning message to the console
// and prepends it with [WARN] in bold yellow text.
func Warn(message ...interface{}) {
	fmt.Println(
		aurora.Yellow(
			aurora.Bold("[WARN]"),
		),
		fmt.Sprint(message...),
	)
}

// Fail prints an error message to the console
// and prepends it with [FAIL] in bold red text.
// It also exits the program with an exit code of 1.
func Fail(message ...interface{}) {
	fmt.Println(
		aurora.Red(
			aurora.Bold("[FAIL]"),
		),
		fmt.Sprint(message...),
	)
	os.Exit(1)
}

// SuccessF prints a successful message given
// a format string to the console and prepends
// it with [SCSS] in bold green text.
func SuccessF(format string, message ...interface{}) {
	fmt.Println(
		aurora.Green(
			aurora.Bold("[SCSS]"),
		),
		fmt.Sprintf(
			format,
			message...,
		),
	)
}

// InfoF prints an informational message given
// a format string to the console and prepends
// it with [SCSS] in bold yellow text.
func InfoF(format string, message ...interface{}) {
	fmt.Println(
		aurora.Cyan(
			aurora.Bold("[INFO]"),
		),
		fmt.Sprintf(
			format,
			message...,
		),
	)
}

// WarnF prints a warning message given
// a format string to the console and prepends
// it with [WARN] in bold yellow text.
func WarnF(format string, message ...interface{}) {
	fmt.Println(
		aurora.Yellow(
			aurora.Bold("[WARN]"),
		),
		fmt.Sprintf(
			format,
			message...,
		),
	)
}

// FailF prints an error message given
// a format string to the console and prepends
// it with [FAIL] in bold red text.
func FailF(format string, message ...interface{}) {
	fmt.Println(
		aurora.Red(
			aurora.Bold("[FAIL]"),
		),
		fmt.Sprintf(
			format,
			message...,
		),
	)
	os.Exit(1)
}
