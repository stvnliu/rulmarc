package main

import (
	"time"

	. "github.com/gbin/goncurses"
	"gitlab.com/stvnliu/ai_game/menu"
	. "gitlab.com/stvnliu/ai_game/utils/types"
	// "gitlab.com/stvnliu/ai_game/utils/window_helper"
)

const (
	INPUT_PROMPT_LENGTH = 40
	INPUT_PROMPT_HEIGHT = 10
)

func InputPrompt(scr *Window) string {
	my, mx := scr.MaxYX()
	w, err := NewWindow(INPUT_PROMPT_HEIGHT, INPUT_PROMPT_LENGTH, (my/2)-(INPUT_PROMPT_HEIGHT/2), (mx/2)-(INPUT_PROMPT_LENGTH/2))
	if err != nil {
		panic("Oh no sth went wrong in input!!")
	}
	w.Box(0, 0)
	Echo(true)
	msg := "Game name: "
  w.MovePrint(0, 1, " New game information ")
	w.MovePrint(2, 2, msg)
	w.Move(2, 2+len(msg))
	input, err := w.GetString(16) // character input box
	if err != nil {
		panic("Oh no sth went wrong in input 2!!")
	}
	w.MovePrint(3, 2, input)
	w.Refresh()
  Echo(false)
	for {
		ch := w.GetChar()
		switch Key(ch) {
		case 'q':
			return input
		}
	}
}
func NewGame(scr *Window) {
	//_, _ := scr.MaxYX()
	game_name := InputPrompt(scr)
	// create new game state
	// println("Creating new game %v...", game_name)
	game := Game{
		SaveGame:  game_name,
		LastSaved: time.Now(),
	}
	my_npcs := MakeNpcs()
	game.DataStored.Npcs = my_npcs
  scr.MovePrintf(1, 2, "Created new game \"%v\"!", game.SaveGame)
  for i:=0; i<len(game.DataStored.Npcs); i++ {
    scr.MovePrintf(2+i, 2, "Initialising \"%v\"...", game.DataStored.Npcs[i].Name)
    scr.MovePrintf(3+i, 2, "Found NPC query string!")
    scr.Refresh()

  }
		// println(game.DataStored.Npcs[0].Name)
}
func Continue(scr *Window) {
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
	}
	menu.CreateMenu(scr, menu_items)
}
