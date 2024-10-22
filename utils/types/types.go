package types

import "time"


type Data struct {
  Npcs []Npc
}
type Game struct {
  SaveGame string 
  LastSaved time.Time 
  DataStored Data
}

func (game Game) NewGame(game_name string, data Data) Game {
  game.SaveGame = game_name
  game.LastSaved = time.Now()
  game.DataStored = data
  return game
}
