package cli

import (
	"GO-GET-SOME-STYLEZ/pkg"
	"fmt"
	"os"
)

func CreateLinuxHelpFile() {
	home, _ := os.UserHomeDir()
	fmt.Print(pkg.Blue + "Creating help file at: " + home + "/oh-my-zsh-setup.txt\n" + pkg.Reset)
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
	content := "1. Run `oh-my-posh config` to choose a theme\n\n" +
		"2. Install Nerd Fonts from https://www.nerdfonts.com\n\n" +
		"3. Set your terminal to use a Nerd Font\n\n" +
		"4. Open a new terminal session and create a new PowerShell profile configuration file running the command:\n\n" +
		"EDIT POSH PROFILE COMMAND:\n" +
		"--------------------------------\n" +
		"notepad $PROFILE\n" +
		"--------------------------------\n" +
		"END OF EDIT POSH PROFILE COMMAND.\n\n" +
		"5. Set the conmfiguration template below copying it and pasting on the notepad window that has just opened:\n\n" +
		"CONFIGURATION TEMPLATE:\n" +
		"--------------------------------\n" +
		"### OH-MY-POSH SETUP ###\n\n" +
		"# winget install JanDeDobbeleer.OhMyPosh -s winget\n\n" +
		"##### POLICY SET #####" +
		"\n\nInstall-Module -Name posh-git -Scope CurrentUser -Force\n" +
		"Install-Module -Name Terminal-Icons -Scope CurrentUser -Force\n" +
		"Install-Module -Name PSReadLine -Scope CurrentUser -Force -SkipPublisherCheck\n\n" +
		"#### GET POSH THEMES ####\n" +
		"Install-Module -Name OhMyPosh.Themes\n\n\n\n" +
		"# REMOVE PROFILE LOADING TIME MESSAGE\n" +
		"#$env:POWERSHELL_TELEMETRY_OPTOUT = 1\n" +
		"#$env:__SuppressAnsiEscapeSequences = $true\n\n" +
		"# Import Modules\n" +
		"Import-Module posh-git\n" +
		"Import-Module Terminal-Icons\n" +
		"Import-Module PSReadLine\n" +
		"Set-PSReadLineOption -HistorySaveStyle SaveIncrementally\n\n" +
		"# Oh My Posh Theme Setup\n" +
		"oh-my-posh init pwsh --config \"C:\\Program Files (x86)\\oh-my-posh\\themes\\darkblood.omp.json\" | Invoke-Expression\n" +
		"--------------------------------\n" +
		"END OF CONFIGURATION TEMPLATE.\n\n" +
		"6. After setting the configuration template. Close any running terminal session and reopen it for a first setup.\n\n" +
		"7. Right after the very first Run using the configuration you just have set. " +
		"Comment the lines below of your PowerShell configuration file to enhance " +
		"the speed of the future terminal sessions you may run:\n\n" +
		"COMMAND LINES TO COMMENT:\n" +
		"--------------------------------\n" +
		"# Install-Module -Name posh-git -Scope CurrentUser -Force\n" +
		"# Install-Module -Name Terminal-Icons -Scope CurrentUser -Force\n" +
		"# Install-Module -Name PSReadLine -Scope CurrentUser -Force -SkipPublisherCheck\n\n" +
		"#### GET POSH THEMES ####\n" +
		"# Install-Module -Name OhMyPosh.Themes" +
		"--------------------------------\n" +
		"END OF EXAMPLE.\n\n" +
		"8. For more Oh-My-Posh themes, visit:\n" +
		"https://ohmyposh.dev/docs/themes\n" +
		"9. END OF SETUP! =D."

	err := os.WriteFile(helpFile, []byte(content), 0644)
	if err != nil {
		fmt.Println(pkg.Red+"Failed to create help file:", err, pkg.Reset)
		return
	}
	fmt.Printf(pkg.Blue+"\nHelp file created at: %s\n"+pkg.Reset, helpFile)
}
