package internal

type DollTypes struct {
	swordDoll DollType
	lanceDoll DollType
	knifeDoll DollType
}

var dollTypes = DollTypes{
	swordDoll: DollType{
		contactDamage: 1,
		texture:       "doll_sword",
	},
	lanceDoll: DollType{
		contactDamage: 2,
		texture:       "doll_lance",
	},
	knifeDoll: DollType{
		contactDamage: 0,
		texture:       "doll_knife",
	},
}
