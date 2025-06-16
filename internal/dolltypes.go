package internal

type DollTypes struct {
	swordDoll DollType
	lanceDoll DollType
}

var dollTypes = DollTypes{
	swordDoll: DollType{
		baseDamage: 1,
		texture:    "doll_sword",
	},
	lanceDoll: DollType{
		baseDamage: 2,
		texture:    "doll_lance",
	},
}
