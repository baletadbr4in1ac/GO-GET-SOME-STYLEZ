package cli

import (
	"GO-GET-SOME-STYLEZ/pkg"
	"fmt"
	"os"
	"os/exec"
)

// BackupZshrc backs up the existing .zshrc file.
func BackupZshrc() {
	home, _ := os.UserHomeDir()
	zshrc := home + "/.zshrc"
	backup := home + "/.zshrc_original_bkp"
	if pkg.Exists(zshrc) {
		err := os.Rename(zshrc, backup)
		if err != nil {
			fmt.Println(pkg.Red+"Failed to backup .zshrc:", err, pkg.Reset)
			//slog.Info(pkg.Red+"Failed to backup .zshrc:", err, pkg.Reset)
			return
		}
		fmt.Printf(pkg.Blue+"\nBacked up .zshrc to: %s\n"+pkg.Reset, backup)
		//slog.Info(pkg.Blue+"\nBacked up .zshrc to: %s\n"+pkg.Reset, backup)

	}
	CreateLinuxHelpFile()
	fmt.Println(pkg.Green + "Done." + pkg.Reset)
	//slog.Info(pkg.Green + "Done." + pkg.Reset)
}

func InstallOhMyZsh() {
	BackupZshrc()
	fmt.Println(pkg.Yellow + "Installing Oh My Zsh..." + pkg.Reset)
	//slog.Info(pkg.Yellow + "Installing Oh My Zsh..." + pkg.Reset)
	cmd := exec.Command("bash", "-c",
		"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(pkg.Red+"Failed to install Oh My Zsh:", err, pkg.Reset)
		//slog.Info(pkg.Red+"Failed to install Oh My Zsh:", err, pkg.Reset)
		return
	}
	fmt.Println(pkg.Green + "Done." + pkg.Reset)
}

func InstallZshPlugins() {
	plugins := []string{"zsh-autosuggestions", "zsh-syntax-highlighting"}
	fmt.Printf(pkg.Yellow+"Installing Oh My Zsh plugins: %s\n"+pkg.Reset, plugins)
	//slog.Info(pkg.Yellow+"Installing Oh My Zsh plugins: %s\n"+pkg.Reset, plugins)
	home, _ := os.UserHomeDir()
	customPluginsDir := home + "/.oh-my-zsh/custom/plugins"

	for _, plugin := range plugins {
		pluginPath := fmt.Sprintf("%s/%s", customPluginsDir, plugin)
		if pkg.Exists(pluginPath) {
			fmt.Printf(pkg.Blue+"Plugin %s already exists, skipping...\n"+pkg.Reset, plugin)
			//slog.Info(pkg.Blue+"Plugin %s already exists, skipping...\n"+pkg.Reset, plugin)
			continue
		}
		cmd := exec.Command("git", "clone",
			fmt.Sprintf("https://github.com/zsh-users/%s.git", plugin), pluginPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(pkg.Red+"Failed to install plugin "+plugin+":", err, pkg.Reset)
			//slog.Error(pkg.Red+"Failed to install plugin "+plugin+":", err, pkg.Reset)
			return
		}
	}
	fmt.Println(pkg.Green + "Done." + pkg.Reset)
	//slog.Info(pkg.Green + "Done." + pkg.Reset)
}

func SourceZshrc() {
	fmt.Println(pkg.Yellow + "Sourcing ~/.zshrc..." + pkg.Reset)
	//slog.Info(pkg.Yellow + "Sourcing ~/.zshrc..." + pkg.Reset)
	home, _ := os.UserHomeDir()
	cmd := exec.Command("zsh", "-c", fmt.Sprintf("source %s/.zshrc", home))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(pkg.Red+"Failed to source ~/.zshrc:", err, pkg.Reset)
		//slog.Error(pkg.Red+"Failed to source ~/.zshrc:", err, pkg.Reset)
		return
	}
	fmt.Println(pkg.Green + "Done." + pkg.Reset)
	//slog.Info(pkg.Green + "Done." + pkg.Reset)
}

func OpenNewZshellSession() {
	fmt.Println(pkg.Yellow + "Opening a new shell for configuration..." + pkg.Reset)
	//slog.Info(pkg.Yellow + "Opening a new shell for configuration..." + pkg.Reset)
	cmd := exec.Command("zsh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(pkg.Red+"Failed to open new shell:", err, pkg.Reset)
		//slog.Error(pkg.Red+"Failed to open new shell:", err, pkg.Reset)
		return
	}
}
