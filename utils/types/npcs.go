package types

type NpcAi struct {
	PromptCharacterString string
	QueryFromTableName    string
}
type Npc struct {
	Name string
	Ai   NpcAi
}


