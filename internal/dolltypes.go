package internal

type DollType struct {
	contactDamage float32
	texture       string
	accel         float32
}

var dollTypes = struct {
	swordDoll DollType
	lanceDoll DollType
	knifeDoll DollType
}{
	swordDoll: DollType{
		contactDamage: 1,
		texture:       "doll_sword",
		accel:         500,
	},
	lanceDoll: DollType{
		contactDamage: 2,
		texture:       "doll_lance",
		accel:         500,
	},
	knifeDoll: DollType{
		contactDamage: 0,
		texture:       "doll_knife",
		accel:         400,
	},
}
