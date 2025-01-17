package types

func InitObjects() {
	WEAPON_OLD_FAMILY_SWORD := Weapon{
    name: "Mjorrsword",
    atk: 10,
    meta: WeaponMetadata{
      info: "The Mjorrsword is an old sword, a legacy in your family. \nLast used in the Fjolrholmer Revolution, it is now yours to hold on to.",
    }
  }
}

type Inventory struct {
	weapons []Weapon
	foods   []Food
	potions []Potion
}
type Player struct {
	name      string
	inventory Inventory
	effects   []Effect
	wallet    []Currency
}

type WeaponMetadata struct {
	info string
}

type Weapon struct {
	name       string
	atk        int
	meta       WeaponMetadata
}

type Food struct {
	name         string
	regen_health int
}
type Effect struct {
	name   string
	effect func(p *Player)
}

type Potion struct {
	name   string
	effect Effect
}

type Currency struct {
	name   string
	prefix string
	value  int
	amount int
}

type Consumable interface {
	Food | Potion
}
