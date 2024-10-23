package main

import (
	. "github.com/gbin/goncurses"
	"gitlab.com/stvnliu/ai_game/menu"
  "gitlab.com/stvnliu/ai_game/utils/windows"
	. "gitlab.com/stvnliu/ai_game/utils/types"
)

const (
	LEFT_PAD  = 3
	RIGHT_PAD = 3
)



func Continue(scr *Window) {
  response := windows.InputPrompt(scr, "Continue!!", "Your answer:", 20)
  scr.MovePrintf(5, 2, "Resp: %v", response)
	// recover state from last save?
}
func Exit(scr *Window) {
	// save game state?
}
func main() {
	scr, err := Init()
	if err != nil {
		println("Something went wrong with Ncurses! Aborting!")
		return
	}
	menu_items := []GameMenuItem{
		{Name: "New game!", Operation: NewGame},
		{Name: "Continue!", Operation: Continue},
		{Name: "Exit!", Operation: Exit},
    {Name: "Test function!", Operation: func(scr *Window) {

    },},
	}
	menu.CreateMenu(scr, menu_items)
}
