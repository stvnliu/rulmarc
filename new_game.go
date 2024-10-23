package main

import (
	"time"

	. "github.com/gbin/goncurses"
	. "gitlab.com/stvnliu/ai_game/utils/types"
	"gitlab.com/stvnliu/ai_game/utils/windows"
  "gitlab.com/stvnliu/ai_game/utils/helper"
)

func NewGame(scr *Window) {
	//_, _ := scr.MaxYX()
	game_name := windows.InputPrompt(scr, " New game information ", "Game name: ", 20)
	// create new game state
	// println("Creating new game %v...", game_name)
	game := Game{
		SaveGame:  game_name,
		LastSaved: time.Now(),
	}
	my_npcs := MakeNpcs()
	game.DataStored.Npcs = my_npcs
	scr.MovePrintf(1, 2, "Created new game \"%v\"!", game.SaveGame)
	for i := 0; i < len(game.DataStored.Npcs); i++ {
		scr.MovePrintf(2+i, 2, "Initialising \"%v\"...", game.DataStored.Npcs[i].Name)
		scr.MovePrintf(3+i, 2, "Found NPC query string!")
		scr.Refresh()

	}
	my, mx := scr.MaxYX()

	// Initialise container box for game content
	w, err := NewWindow(my-1, mx-1, 1, 1)
	if err != nil {
		panic("Oh shit something happened that shouldn't")
	}
	w.Box(0, 0)
  input_window, input_window_error := NewWindow(6,mx-3,my-7,2)
	if input_window_error != nil {
    panic("Oh no")
  }
  input_window.Box(1, 1)
  input_window.MovePrint(1, 1, "> ")
  input_window.Move(1, 3)
  w.Refresh()
  input_window.Refresh()
  texts := []string {
    "Hello world!!",
    "Welcome to R.U.L.M.A.R.C.",
    "This is an experimental game project that uses Large Language Models to power realistically rendered characters.",
    "============ Copyright 2024 @ Zhongheng Liu & Zhiyong He =============",
    "Try it! Put in some characters in the input box below!",
    "Please wait while we boot some Artificial Intelligencce models for the first part of the game...",
  }
  for i := 0; i < len(texts); i++ {
    helper.IncrementalPrint(w, texts[i], 1+i, 1, 500)
  }
	for {
		ch := w.GetChar()
		switch Key(ch) {
		case 'q':
			w.Erase()
			w.Refresh()
			return
		}
		w.Refresh()
	}
	// println(game.DataStored.Npcs[0].Name)
}
