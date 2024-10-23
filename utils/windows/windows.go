package windows

import (
  . "github.com/gbin/goncurses"
)
const (
  LEFT_PAD = 2 
  RIGHT_PAD = 2
)
func InputPrompt(
	scr *Window,
	title string,
	prompt string,
	input_length int,
) string {
	height, length := 8, len(prompt)+input_length+LEFT_PAD+RIGHT_PAD
	my, mx := scr.MaxYX()
	w, err := NewWindow(
		height,
		length,
		(my/2)-(height/2),
		(mx/2)-(length/2),
	)
	if err != nil {
		panic("Oh no sth went wrong in input!!")
	}
	w.Box(0, 0)
	Echo(true)
	w.MovePrint(0, 1, title)
	w.MovePrint(2, 2, prompt)
	w.Move(2, 2+len(prompt))
	input, err := w.GetString(input_length) // character input box
	if err != nil {
		panic("Oh no sth went wrong in input 2!!")
	}
	w.MovePrintf(height-2, 2, "Press q to continue...")
	w.Refresh()
	Echo(false)
	for {
		ch := w.GetChar()
		switch Key(ch) {
		case 'q':
			w.Erase()
			w.Refresh()
			w.Delete()
			return input
		}
	}
}
