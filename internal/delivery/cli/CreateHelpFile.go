package cli

import (
	"GO-GET-SOME-STYLEZ/pkg"
	"fmt"
	"os"
)

func CreateLinuxHelpFile() {
	home, _ := os.UserHomeDir()
	fmt.Print(pkg.Blue + "Creating help file at: " + home + "/oh-my-zsh-setup.txt\n" + pkg.Reset)
	//slog.Info(pkg.Blue + "Creating help file at: " + home + "/oh-my-zsh-setup.txt\n" + pkg.Reset)
	helpFile := home + "/oh-my-zsh-setup.txt"
	content := "1. For the plugins to work correctly.\nMake sure you include the following plugin configuration in your " +
		"~/.zshrc when the installation process is finished\n plugins=(git zsh-autosuggestions zsh-syntax-highlighting)\n\n" +
		"2. Install Nerd Fonts from https://www.nerdfonts.com\n\n" +
		"3. Set your terminal to use a Nerd Font\n"
	err := os.WriteFile(helpFile, []byte(content), 0644)
	if err != nil {
		fmt.Println(pkg.Red+"Failed to create help file:", err, pkg.Reset)
		return
	}
	fmt.Printf(pkg.Blue+"\nHelp file created at: %s\n"+pkg.Reset, helpFile)
	//slog.Info(pkg.Blue+"\nHelp file created at: %s\n"+pkg.Reset, helpFile)
}
func CreateWindowsHelpFile() {
	desktop, _ := os.UserHomeDir()
	helpFile := desktop + "\\Desktop\\oh-my-posh-setup.txt"
	content := "1. Run `oh-my-posh config` to choose a theme\n" +
		"2. Install Nerd Fonts from https://www.nerdfonts.com\n" +
		"3. Set your terminal to use a Nerd Font\n"
	err := os.WriteFile(helpFile, []byte(content), 0644)
	if err != nil {
		fmt.Println(pkg.Red+"Failed to create help file:", err, pkg.Reset)
		return
	}
	fmt.Printf(pkg.Blue+"\nHelp file created at: %s\n"+pkg.Reset, helpFile)
	//slog.Info(pkg.Blue+"\nHelp file created at: %s\n"+pkg.Reset, helpFile)
}
