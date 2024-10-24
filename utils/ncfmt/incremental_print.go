package ncfmt

import (
	"time"

	. "github.com/gbin/goncurses"
)

func IncrementalPrintMany(
	w *Window,
	y int,
	x int,
	texts []string,
	duration time.Duration,
) {
	for i := 0; i < len(texts); i++ {
		IncrementalPrint(w, texts[i], y+i, x, int(1*time.Second))
	}
}
func IncrementalPrint(scr *Window, text string, from_y int, from_x int, interval_millis int) {
	for i := 0; i < len(text); i++ {
		ch := string([]rune(text)[i])
		_, mx := scr.MaxYX()
		cy := i/mx + from_y
		cx := i%mx + 1
		scr.MovePrint(cy, cx, ch)
		time.Sleep(time.Duration(1000/len(text)) * time.Millisecond)
		scr.Refresh()
	}
}
