package pkg

import (
	"os"
	"os/exec"
)

func ClearConsole() string {
	osConsoleToClean := OSDetectionSimple()
	switch osConsoleToClean {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return ""
		}
	case "windows":
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return ""
		}
		return "windows"
	}
	return "unknown"
}
