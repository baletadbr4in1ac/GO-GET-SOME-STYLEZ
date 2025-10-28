package cli

import (
	"GO-GET-SOME-STYLEZ/pkg"
	"fmt"
	"os"
	"os/exec"
)

const ohmyposhCommand = "oh-my-posh"

func InstallOhMyPosh() {
	isOhmyposhInstalled := pkg.IsItInstalled(ohmyposhCommand)

	switch isOhmyposhInstalled {
	case true:
		fmt.Println(pkg.Yellow + "Oh My Posh is already installed." + pkg.Reset)
		fmt.Println(pkg.Yellow + "Now Upgrading the installed version instead..." + pkg.Reset)
		//slog.Info(pkg.Yellow + "Installing Oh My Posh for Windows..." + pkg.Reset)
		cmd := exec.Command("winget", "install", "JanDeDobbeleer.OhMyPosh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//err := cmd.Run()
		cmd.Run()

		CreateWindowsHelpFile()
		fmt.Println(pkg.Purple + "Done setting up Oh My Posh!" + pkg.Reset)
		return
	case false:
		fmt.Println(pkg.Yellow + "Installing Oh My Posh for Windows..." + pkg.Reset)
		//slog.Info(pkg.Yellow + "Installing Oh My Posh for Windows..." + pkg.Reset)
		cmd := exec.Command("winget", "install", "JanDeDobbeleer.OhMyPosh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			fmt.Println(pkg.Red+"Failed during main installation process:", err, pkg.Reset)
			//slog.Error(pkg.Red+"Failed to install Oh My Posh:", err, pkg.Reset)
			return
		}
	}
	return
}
