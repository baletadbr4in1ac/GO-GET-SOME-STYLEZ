package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallPackage(manager, packageToInstall string, OpSystem string) {
	fmt.Printf(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)

	switch manager {
	case "fedora":
		RunInstallPkgCommand(manager, packageToInstall, OpSystem)
		return
	case "debian":
		RunInstallPkgCommand(manager, packageToInstall, OpSystem)
		return
	case "arch":
		RunInstallPkgCommand(manager, packageToInstall, OpSystem)
		return
	case "manjaro":
		RunInstallPkgCommand(manager, packageToInstall, OpSystem)
		return
	case "opensuse":
		RunInstallPkgCommand(manager, packageToInstall, OpSystem)
		return
	case "darwin":
		RunInstallPkgCommand(manager, packageToInstall, OpSystem)
		return
	case "windows":
		RunInstallPkgCommand(manager, packageToInstall, OpSystem)
		return

	}
}

func RunInstallPkgCommand(manager, packageToInstall string, OpSystem string) {

	switch OpSystem {
	case "fedora":
		//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)
		cmd := exec.Command("sudo", manager, "install", "-y", packageToInstall)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(Red+"Failed to install "+packageToInstall+":", err, Reset)
			//slog.Error(Red+"Failed to install "+packageToInstall+":", err, Reset)
			return
		}
		fmt.Println(Green + "Done." + Reset)
		//slog.Info(Green + "Done." + Reset)
		return
	case "debian":
		//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)
		cmd := exec.Command("sudo", manager, "install", "-y", packageToInstall)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(Red+"Failed to install "+packageToInstall+":", err, Reset)
			//slog.Error(Red+"Failed to install "+packageToInstall+":", err, Reset)
			return
		}
		fmt.Println(Green + "Done." + Reset)
		//slog.Info(Green + "Done." + Reset)
		return
	case "arch":
		//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)
		cmd := exec.Command("sudo", manager, "install", "--noconfirm", packageToInstall)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(Red+"Failed to install "+packageToInstall+":", err, Reset)
			//slog.Error(Red+"Failed to install "+packageToInstall+":", err, Reset)
			return
		}
		fmt.Println(Green + "Done." + Reset)
		//slog.Info(Green + "Done." + Reset)
		return
	case "manjaro":
		//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)
		cmd := exec.Command("sudo", manager, "install", "-y", packageToInstall)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(Red+"Failed to install "+packageToInstall+":", err, Reset)
			//slog.Error(Red+"Failed to install "+packageToInstall+":", err, Reset)
			return
		}
		fmt.Println(Green + "Done." + Reset)
		//slog.Info(Green + "Done." + Reset)
		return
	case "opensuse":
		//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)
		cmd := exec.Command("sudo", manager, "install", "-n", packageToInstall)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(Red+"Failed to install "+packageToInstall+":", err, Reset)
			//slog.Error(Red+"Failed to install "+packageToInstall+":", err, Reset)
			return
		}
		fmt.Println(Green + "Done." + Reset)
		//slog.Info(Green + "Done." + Reset)
		return
	case "darwin":
		//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)
		cmd := exec.Command(manager, "install", packageToInstall)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(Red+"Failed to install "+packageToInstall+":", err, Reset)
			//slog.Error(Red+"Failed to install "+packageToInstall+":", err, Reset)
			return
		}
		fmt.Println(Green + "Done." + Reset)
		//slog.Info(Green + "Done." + Reset)
		return
	case "windows":
		//slog.Info(Yellow+"Installing %s using %s...\n"+Reset, packageToInstall, manager)
		cmd := exec.Command(manager, "install", packageToInstall)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(Red+"Failed to install "+packageToInstall+":", err, Reset)
			//slog.Error(Red+"Failed to install "+packageToInstall+":", err, Reset)
			return
		}
		fmt.Println(Green + "Done." + Reset)
		//slog.Info(Green + "Done." + Reset)
	}

}
