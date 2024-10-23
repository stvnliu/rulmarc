package main

import (
	. "gitlab.com/stvnliu/ai_game/utils/types"
)

func MakeNpcs() []Npc {
	npcs := []Npc{}
	helper01 := Npc{
		Name: "Helper01_NPC",
		Ai: NpcAi{
			PromptCharacterString: "You are a new helper assisting new players of a role-playing game set in $SCENE$, in a village called $VILLAGE$. With the information immediately preceeding, output only what you would say to a new player who just arrived in the village to provide helpful guidance.",
			QueryFromTableName:    "helper",
		},
	}
	npcs = append(npcs, helper01)

	rulmarc := Npc{
		Name: "Rulmarc",
		Ai: NpcAi{
			PromptCharacterString: "You are a medieval villager called Rulmarc.",
			QueryFromTableName:    "helper",
		},
	}

	npcs = append(npcs, rulmarc)

	return npcs
}
