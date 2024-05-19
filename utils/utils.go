package utils

import (
	"os"
	"os/exec"
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
