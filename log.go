package stat

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func Success(message string) {
	fmt.Println(aurora.Green(aurora.Bold("[SCSS]")), message)
}

func Info(message string) {
	fmt.Println(aurora.Cyan(aurora.Bold("[INFO]")), message)
}

func Warn(message string) {
	fmt.Println(aurora.Yellow(aurora.Bold("[WARN]")), message)
}

func Fatal(message string) {
	fmt.Println(aurora.Red(aurora.Bold("[FATL]")), message)
	os.Exit(1)
}
