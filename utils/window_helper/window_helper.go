package window_helper
import (
  . "github.com/gbin/goncurses"
)

func CreateMenu(win *Window, menu_string []string) {
  x, y := 2, 2 
  win.Clear()
  win.Box(0, 0)
  for i, str := range menu_string {
    win.MovePrint(y+i, x, str)
  }
}
