package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

func main() {
	fmt.Println(blue + "\nStarting GO-GET-SOME-STYLEZ..." + reset)
	backupZshrc()
	switch detectOS() {
	case "fedora":
		installPackage("dnf", "zsh")
	case "debian", "ubuntu":
		installPackage("apt", "zsh")
	case "arch", "manjaro":
		installPackage("pacman", "zsh")
	case "opensuse":
		installPackage("zypper", "zsh")
	case "darwin":
		installPackage("brew", "zsh")
	case "windows":
		installOhMyPosh()
		return
	default:
		fmt.Println(red + "Unsupported OS" + reset)
		return
	}

	installOhMyZsh()
	installZshPlugins()

	if askDefaultShell() {
		setDefaultShell()
	}

	openNewShell()
	fmt.Println(green + "\nNOW YOU GOT STYLEZ 😉" + reset)
}

func detectOS() string {
	switch runtime.GOOS {
	case "linux":
		if exists("/etc/fedora-release") {
			return "fedora"
		} else if exists("/etc/debian_version") {
			return "debian"
		} else if exists("/etc/arch-release") {
			return "arch"
		} else if exists("/etc/manjaro-release") {
			return "manjaro"
		} else if exists("/etc/os-release") {
			return "opensuse"
		}
	case "darwin":
		return "darwin"
	case "windows":
		return "windows"
	}
	return "unknown"
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func installPackage(manager, pkg string) {
	fmt.Printf(yellow+"Installing %s using %s...\n"+reset, pkg, manager)
	err := exec.Command("sudo", manager, "install", "-y", pkg).Run()
	if err != nil {
		return
	}
	fmt.Println(green + "Done." + reset)
}

func installOhMyZsh() {
	fmt.Println(yellow + "Installing Oh My Zsh..." + reset)
	err := exec.Command("sh", "-c",
		"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)").Run()
	if err != nil {
		return
	}
	fmt.Println(green + "Done." + reset)
}

func installZshPlugins() {
	plugins := []string{"zsh-autosuggestions", "zsh-syntax-highlighting"}
	fmt.Printf(yellow+"Installing Oh My Zsh plugins: %s\n"+reset, plugins)
	for _, plugin := range plugins {
		err := exec.Command("git", "clone", fmt.Sprintf("https://github.com/zsh-users/%s.git", plugin),
			fmt.Sprintf("$ZSH_CUSTOM/plugins/%s", plugin)).Run()
		if err != nil {
			return
		}
	}
	fmt.Println(green + "Done." + reset)
}

func askDefaultShell() bool {
	fmt.Print(blue + "Set Zsh as default shell? (Y/N): " + reset)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input == "Y\n" || input == "y\n"
}

func setDefaultShell() {
	fmt.Println(yellow + "Setting Zsh as default shell..." + reset)
	err := exec.Command("chsh", "-s", "/bin/zsh").Run()
	if err != nil {
		return
	}
	fmt.Println(green + "\nDone." + reset)
}

func backupZshrc() {
	home, _ := os.UserHomeDir()
	zshrc := home + "/.zshrc"
	backup := home + "/.zshrc_original_bkp"
	if exists(zshrc) {
		err := os.Rename(zshrc, backup)
		if err != nil {
			return
		}
		fmt.Printf(blue+"\nBacked up .zshrc to: %s\n"+reset, backup)
	}
}

func installOhMyPosh() {
	fmt.Println(yellow + "Installing Oh My Posh for Windows..." + reset)
	err := exec.Command("winget", "install", "JanDeDobbeleer.OhMyPosh").Run()
	if err != nil {
		return
	}
	fmt.Println(green + "\nDone." + reset)
	createWindowsHelpFile()
}

func createWindowsHelpFile() {
	desktop, _ := os.UserHomeDir()
	helpFile := desktop + "\\Desktop\\oh-my-posh-setup.txt"
	content := "1. Run `oh-my-posh config` to choose a theme\n" +
		"2. Install Nerd Fonts from https://www.nerdfonts.com\n" +
		"3. Set your terminal to use a Nerd Font\n"
	err := os.WriteFile(helpFile, []byte(content), 0644)
	if err != nil {
		return
	}
	fmt.Printf(blue+"\nHelp file created at: %s\n"+reset, helpFile)
}

func openNewShell() {
	fmt.Println(yellow + "Opening a new shell for configuration..." + reset)
	err := exec.Command("zsh").Run()
	if err != nil {
		return
	}
}
