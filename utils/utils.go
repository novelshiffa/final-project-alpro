package utils

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearTerminal membersihkan terminal konsol dengan menggunakan perintah sistem operasi yang sesuai.
// Untuk Windows, perintah "cls" digunakan, sedangkan untuk sistem Unix-like, perintah "clear" digunakan.
func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // untuk Windows
	} else {
		cmd = exec.Command("clear") // untuk sistem Unix-like
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
