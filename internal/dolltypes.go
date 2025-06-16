package internal

type DollTypes struct {
	knifeDoll DollType
	lanceDoll DollType
}

var dollTypes = DollTypes{
	knifeDoll: DollType{
		baseDamage: 1,
		texture:    "doll",
	},
	lanceDoll: DollType{
		baseDamage: 2,
		texture:    "doll_lance",
	},
}
