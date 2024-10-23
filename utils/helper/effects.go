package helper

import (
	"fmt"
	"time"

	. "github.com/gbin/goncurses"
)

func BlinkCursorUntilInput(scr *Window, pos_y int, pos_x int, interval time.Duration) Key {
	scr.Move(pos_y, pos_x)
	var activation_key Key
	for {
		Cursor(2)
		scr.Timeout(int((interval / 3).Milliseconds()))
		activation_key = scr.GetChar()
		if activation_key != 0 {
			break
		}
		time.Sleep(interval / 3)
		Cursor(0)
		time.Sleep(interval / 3)
	}
	return activation_key
}
func BlinkCursorUntilDone(scr *Window, pos_y int, pos_x int, interval time.Duration, done <-chan bool) {
	scr.Move(pos_y, pos_x)
	for {
		Cursor(2)
		select {
		case is_done, ok := <-done:
			if ok && is_done {
				return
			} else {
				fmt.Println("Channel closed?")
			}
		default:
			time.Sleep(interval / 2)
			Cursor(0)
			time.Sleep(interval / 2)

		}
	}
}
func BlinkCursorWithTime(scr *Window, pos_y int, pos_x int, duration time.Duration, interval time.Duration) {
	scr.Move(pos_y, pos_x)
	n := duration / interval
	for i := 0; i < int(n); i++ {
		Cursor(2)
		time.Sleep(interval / 2)
		Cursor(0)
		time.Sleep(interval / 2)
	}
}
