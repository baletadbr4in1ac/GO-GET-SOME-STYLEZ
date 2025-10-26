package pkg

import (
	"github.com/common-nighthawk/go-figure"
)

func PrintMainBanner() {
	mainBannerPartOne := figure.NewColorFigure("Go Get", "elite", "green", true)
	mainBannerPartTwo := figure.NewColorFigure("   Some Stylez", "elite", "green", true)

	ClearConsole()

	mainBannerPartOne.Print()
	mainBannerPartTwo.Print()

}
