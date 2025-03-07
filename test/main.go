package main

import (
	"webscr/core"

	"github.com/fatih/color"
)

func main() {
	ColorPurplePrint := color.New()
	ColorPurplePrint.Add(color.FgHiRed)

	ColorYellowPrint := color.New()
	ColorYellowPrint.Add(color.FgHiYellow)

	ColorPurplePrint.Println(`
    _       __     __   _____          
   | |     / /__  / /_ / ___/__________
   | | /| / / _ \/ __ \\__ \/ ___/ ___/
   | |/ |/ /  __/ /_/ /__/ / /__/ /    
   |__/|__/\___/_.___/____/\___/_/ v2.0.1
		   `)
	ColorYellowPrint.Println("-h, --help 查看帮助")

	core.Runscreenshot()
}
