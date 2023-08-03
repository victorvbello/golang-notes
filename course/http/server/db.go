package server

type HeroSkill struct {
	Kind   string
	Damage float32
	Energy int32
}

type Hero struct {
	ID     int
	Name   string
	Alias  string
	Skills []HeroSkill
}

var DB = struct {
	Heroes []Hero
}{
	Heroes: []Hero{
		{
			ID:    1,
			Alias: "Darth vader",
			Name:  "Anakin Skywalker",
			Skills: []HeroSkill{
				{Kind: "Lightsaber", Damage: 10.5, Energy: 50},
				{Kind: "Force Healing", Damage: 0.0, Energy: 400},
				{Kind: "Force Lightning", Damage: 50.10, Energy: 300},
			},
		},
		{
			ID:    2,
			Alias: "Mando",
			Name:  "Din Djarin",
			Skills: []HeroSkill{
				{Kind: "Blaster", Damage: 100.10, Energy: 100},
				{Kind: "Hand-to-Hand combat", Damage: 50.0, Energy: 300},
			},
		},
	},
}
