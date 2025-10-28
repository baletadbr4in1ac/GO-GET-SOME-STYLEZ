package pkg

import (
	"os"
	"os/exec"
	"runtime"
)

func OSDetectionDetailed() string {
	switch runtime.GOOS {
	case "linux":
		if Exists("/etc/fedora-release") {
			return "fedora"
		} else if Exists("/etc/debian_version") {
			return "debian"
		} else if Exists("/etc/arch-release") {
			return "arch"
		} else if Exists("/etc/manjaro-release") {
			return "manjaro"
		} else if Exists("/etc/os-release") {
			return "opensuse"
		}
	case "darwin":
		return "darwin"
	case "windows":
		return "windows"
	}
	return "unknown"
}

func OSDetectionSimple() string {
	detectedOs := runtime.GOOS
	return detectedOs
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsItInstalled(toolCmd string) bool {
	_, err := exec.LookPath(toolCmd)
	return err == nil
}
