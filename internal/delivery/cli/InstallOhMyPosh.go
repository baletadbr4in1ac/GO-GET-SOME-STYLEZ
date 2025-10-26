package cli

import (
	"GO-GET-SOME-STYLEZ/pkg"
	"fmt"
	"os"
	"os/exec"
)

func InstallOhMyPosh() {
	fmt.Println(pkg.Yellow + "Installing Oh My Posh for Windows..." + pkg.Reset)
	//slog.Info(pkg.Yellow + "Installing Oh My Posh for Windows..." + pkg.Reset)
	cmd := exec.Command("winget", "install", "JanDeDobbeleer.OhMyPosh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(pkg.Red+"Failed to install Oh My Posh:", err, pkg.Reset)
		//slog.Error(pkg.Red+"Failed to install Oh My Posh:", err, pkg.Reset)
		return
	}
	CreateWindowsHelpFile()
	fmt.Println(pkg.Green + "\nDone." + pkg.Reset)
	//slog.Info(pkg.Green + "\nDone." + pkg.Reset)
}
