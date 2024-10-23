package main

import (
	// "fmt"
	"time"

	. "github.com/gbin/goncurses"
	"gitlab.com/stvnliu/ai_game/utils/helper"
	. "gitlab.com/stvnliu/ai_game/utils/types"
	"gitlab.com/stvnliu/ai_game/utils/windows"
)

const (
	STD_BLINK_INTERVAL = 450 * time.Millisecond
)

func IncrementalPrintMany(
	w *Window,
	y int,
	x int,
	texts []string,
	duration time.Duration,
) {
	for i := 0; i < len(texts); i++ {
		helper.IncrementalPrint(w, texts[i], y+i, x, int(1*time.Second))
	}
}
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
	input_window, input_window_error := NewWindow(6, mx-3, my-7, 2)
	if input_window_error != nil {
		panic("Oh no")
	}
	input_window.Box(1, 1)
	input_window.MovePrint(1, 1, "> ")
	input_window.Move(1, 3)
	w.Refresh()
	input_window.Refresh()
	texts := []string{
		"Hello world!!",
		"Welcome to R.U.L.M.A.R.C.",
		"This is an experimental game project that uses Large Language Models to power realistically rendered characters.",
		"============ Copyright 2024 @ Zhongheng Liu & Zhiyong He =============",
		"Please wait while we boot some AI models for the first part of the game...",
	}
	IncrementalPrintMany(w, 1, 1, texts, time.Duration(1*time.Second))

	init_done := make(chan bool, 1)

	go helper.BlinkCursorUntilDone(
		w,
		len(texts)+1,
		1,
		STD_BLINK_INTERVAL,
		init_done,
	)

	// Simulating game init process
	// time.Sleep(time.Duration(10 * time.Second))
	init_done <- true // can trigger blinker process finish
	texts2 := []string{
		"Ok we are done with everything!",
		"Now try putting something in the input box below!",
	}
	IncrementalPrintMany(w, len(texts)+1, 1, texts2, time.Duration(1*time.Second))
	key := helper.BlinkCursorUntilInput(input_window, 1, 3, STD_BLINK_INTERVAL)
	Cursor(0)
  my_input := "You said: "
	for {
		Cursor(2)
		_, cx := input_window.CursorYX()
		if key != 0 {
			// workaround for backspace key
			if (key == KEY_BACKSPACE) || (key == 127) {
				if cx-1 > 2 {
					input_window.MoveDelChar(1, cx-1)
          my_input = my_input[:len(my_input)-1]
				}
			} else if !((key == KEY_ENTER) || (key == KEY_RETURN)) {
				input_window.MovePrint(1, cx, KeyString(key))
        my_input += KeyString(key)
				input_window.Move(1, cx+1)
			} else {
				break
			}
		}
		key = input_window.GetChar()
	}

  helper.IncrementalPrint(w, my_input, 8, 1, int(time.Duration(1*time.Second)))
	// User input processing
	for {
		ch := w.GetChar()
		switch Key(ch) {
		case 'q':
			w.Erase()
			w.Refresh()
			w.Delete()
			return
		}
		w.Refresh()
	}
	// println(game.DataStored.Npcs[0].Name)
}
