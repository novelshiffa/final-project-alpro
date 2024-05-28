package utils

import (
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
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

func IsNumericString(input string) bool {
	pattern := `^\d+$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(input)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
