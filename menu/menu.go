package menu

import (
	. "github.com/gbin/goncurses"
	. "gitlab.com/stvnliu/ai_game/utils/types"
	// "log"
)

const (
	HEIGHT = 10
	WIDTH  = 30
)

func CreateMenu(stdscr *Window, menu []GameMenuItem) {
	var active int
	defer End()

	Raw(true)
	Echo(false)
	Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	_, mx := stdscr.MaxYX()
	y, x := 2, (mx/2)-(WIDTH/2)

	win, _ := NewWindow(HEIGHT, WIDTH, y, x)
	win.Keypad(true)

	stdscr.Print("Welcome to Rulmarc, the Role-playing game using LLMs (Use q to quit the current window)")
	stdscr.Refresh()

	printmenu(win, menu, active)

	for {
		ch := stdscr.GetChar()
		switch Key(ch) {
		case 'q':
			win.Erase()
			win.Refresh()
			return
		case KEY_UP:
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}
		case KEY_DOWN:
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}
		case KEY_RETURN, KEY_ENTER, Key('\r'):
			menu[active].Operation(stdscr)
			stdscr.ClearToEOL()
			win.Erase()
			win.Refresh()
			stdscr.Refresh()
			return
		default:
			stdscr.ClearToEOL()
			stdscr.Refresh()
		}

		printmenu(win, menu, active)
	}
}

func printmenu(w *Window, menu []GameMenuItem, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	for i, item := range menu {
		if i == active {
			w.AttrOn(A_REVERSE)
			w.MovePrint(y+i, x, item.Name)
			w.AttrOff(A_REVERSE)
		} else {
			w.MovePrint(y+i, x, item.Name)
		}
	}
	w.Refresh()
}
