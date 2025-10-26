package cli

import (
	"GO-GET-SOME-STYLEZ/pkg"
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func GetDefaultShell() bool {
	fmt.Print(pkg.Blue + "Set Zsh as default shell? (Y/N): " + pkg.Reset)
	//slog.Info(pkg.Blue + "Set Zsh as default shell? (Y/N): " + pkg.Reset)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input == "Y\n" || input == "y\n"
}

func SetDefaultShell() {
	fmt.Println(pkg.Yellow + "Setting Zsh as default shell..." + pkg.Reset)
	//slog.Info(pkg.Yellow + "Setting Zsh as default shell..." + pkg.Reset)
	cmd := exec.Command("chsh", "-s", "/bin/zsh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(pkg.Red+"Failed to set Zsh as default shell:", err, pkg.Reset)
		return
	}
	fmt.Println(pkg.Green + "\nDone." + pkg.Reset)
	//slog.Info(pkg.Green + "\nDone." + pkg.Reset)
}
