package main

import (
  "gitlab.com/stvnliu/ai_game/menu"
)

func main() {
  menu_items := []string{"New game", "Continue", "Exit"}
  menu.CreateMenu(menu_items)
}
