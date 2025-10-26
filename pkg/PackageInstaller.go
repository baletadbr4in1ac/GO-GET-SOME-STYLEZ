package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallPackage(manager, pachage string) {
	fmt.Printf(Yellow+"Installing %s using %s...\n"+Reset, pachage, manager)
	//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, pachage, manager)
	cmd := exec.Command("sudo", manager, "install", "-y", pachage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(Red+"Failed to install "+pachage+":", err, Reset)
		//slog.Error(Red+"Failed to install "+pachage+":", err, Reset)
		return
	}
	fmt.Println(Green + "Done." + Reset)
	//slog.Info(Green + "Done." + Reset)
}
