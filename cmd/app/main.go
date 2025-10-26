package main

import (
	"GO-GET-SOME-STYLEZ/internal/delivery/cli"
	"GO-GET-SOME-STYLEZ/pkg"
	"fmt"
)

func main() {
	pkg.PrintMainBanner()
	//logger.InitLogger()

	fmt.Println(pkg.Blue + "\n\nStarting GO-GET-SOME-STYLEZ...`\n" + pkg.Reset)
	//slog.Info(pkg.Blue + "\n\nStarting GO-GET-SOME-STYLEZ...`\n" + pkg.Reset)

	switch pkg.OSDetectionDetailed() {
	case "fedora":
		pkg.InstallPackage("dnf", "zsh")
		cli.InstallOhMyZsh()
		cli.InstallZshPlugins()
		if cli.GetDefaultShell() {
			cli.SetDefaultShell()
		}
		cli.SourceZshrc()
		cli.OpenNewZshellSession()
		fmt.Println(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		//slog.Info(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		return
	case "debian", "ubuntu":
		pkg.InstallPackage("apt", "zsh")
		cli.InstallOhMyZsh()
		cli.InstallZshPlugins()
		if cli.GetDefaultShell() {
			cli.SetDefaultShell()
		}
		cli.SourceZshrc()
		cli.OpenNewZshellSession()
		fmt.Println(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		//slog.Info(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		return
	case "arch", "manjaro":
		pkg.InstallPackage("pacman", "zsh")
		cli.InstallOhMyZsh()
		cli.InstallZshPlugins()
		if cli.GetDefaultShell() {
			cli.SetDefaultShell()
		}
		cli.SourceZshrc()
		cli.OpenNewZshellSession()
		fmt.Println(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		//slog.Info(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		return
	case "opensuse":
		pkg.InstallPackage("zypper", "zsh")
		cli.InstallOhMyZsh()
		cli.InstallZshPlugins()
		if cli.GetDefaultShell() {
			cli.SetDefaultShell()
		}
		cli.SourceZshrc()
		cli.OpenNewZshellSession()
		fmt.Println(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		//slog.Info(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		return
	case "darwin":
		pkg.InstallPackage("brew", "zsh")
		cli.InstallOhMyZsh()
		cli.InstallZshPlugins()
		if cli.GetDefaultShell() {
			cli.SetDefaultShell()
		}
		cli.SourceZshrc()
		cli.OpenNewZshellSession()
		fmt.Println(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		//slog.Info(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		return
	case "windows":
		cli.InstallOhMyPosh()
		fmt.Println(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		//slog.Info(pkg.Green + "\nNOW YOU GOT STYLEZ 😉\n" + pkg.Reset)
		return
	default:
		fmt.Println(pkg.Red + "Unsupported OS" + pkg.Reset)
		return
	}
}
