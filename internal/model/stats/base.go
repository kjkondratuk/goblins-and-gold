package stats

type BaseStats struct {
	Str int
	Dex int
	Con int
	Int int
	Wis int
	Cha int
}

type modifiers struct {
	Str modifier
	Dex modifier
	Con modifier
	Int modifier
	Wis modifier
	Cha modifier
}
