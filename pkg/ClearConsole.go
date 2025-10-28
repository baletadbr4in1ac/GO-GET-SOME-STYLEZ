package pkg

import (
	"fmt"
)

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
	return
}

//func ClearConsole() string {
//	osConsoleToClean := OSDetectionSimple()
//	switch osConsoleToClean {
//	case "linux":
//		cmd := exec.Command("clear")
//		cmd.Stdout = os.Stdout
//		err := cmd.Run()
//		if err != nil {
//			return ""
//		}
//	case "windows":
//		cmd := exec.Command("cls")
//		cmd.Stdout = os.Stdout
//		err := cmd.Run()
//		if err != nil {
//			return ""
//		}
//		return "windows"
//	}
//	return "unknown"
//}
