package types

import (
  . "github.com/gbin/goncurses"
)

type GameMenuItem struct {
  Name string
  Operation func(*Window)
}
