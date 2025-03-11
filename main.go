package main

import (
	"webscr/common"
	"webscr/core"
)

func main() {
	version := "2.0.5"

	colorRedPrint := common.Colors(common.ColorRed)

	colorYellowPrint := common.Colors(common.ColorYellow)

	colorRedPrint.Println(`
    _       __     __   _____          
   | |     / /__  / /_ / ___/__________
   | | /| / / _ \/ __ \\__ \/ ___/ ___/
   | |/ |/ /  __/ /_/ /__/ / /__/ /    
   |__/|__/\___/_.___/____/\___/_/ v` + version + `
		   `)
	colorYellowPrint.Println("-h, --help 查看帮助")

	core.Runscreenshot()
}
