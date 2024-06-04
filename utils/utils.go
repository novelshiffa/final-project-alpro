package utils

import (
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	} else {
		cmd = exec.Command("clear") // for Unix-like systems
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func IsValidDateTimeFormat(dateTime string) bool {
	// Define the regex pattern for the datetime format
	// ^ - start of the string
	// \d{4} - exactly 4 digits for the year
	// - - hyphen separator
	// \d{2} - exactly 2 digits for the month
	// - - hyphen separator
	// \d{2} - exactly 2 digits for the day
	// \s - space separator
	// \d{2} - exactly 2 digits for the hour
	// : - colon separator
	// \d{2} - exactly 2 digits for the minute
	// : - colon separator
	// \d{2} - exactly 2 digits for the second
	// $ - end of the string
	pattern := `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Match the input string against the regex pattern
	return re.MatchString(dateTime)
}
