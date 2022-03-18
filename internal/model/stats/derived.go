package stats

const (
	P_NO_PROF = iota
	P_PROF
	P_EXP
)

type DerivedStats struct {
	Mod    modifiers
	Init   int
	AC     int
	SDC    int
	Skills skills
	Saves  saves
}

type modifiers struct {
	Str modifier
	Dex modifier
	Con modifier
	Int modifier
	Wis modifier
	Cha modifier
}

type modifier struct {
	Bonus int
	Prof  Proficiency
}

type Proficiency int

type skills struct {
	Acrobatics     modifier
	AnimalHandling modifier
	Arcana         modifier
	Athletics      modifier
	Deception      modifier
	History        modifier
	Insight        modifier
	Intimidation   modifier
	Investigation  modifier
	Medicine       modifier
	Nature         modifier
	Perception     modifier
	Performance    modifier
	Persuasion     modifier
	Religion       modifier
	SlightOfHand   modifier
	Stealth        modifier
	Survival       modifier
}

type saves struct {
	Str modifier
	Dex modifier
	Con modifier
	Int modifier
	Wis modifier
	Cha modifier
}

func statify(stat int) int {
	return stat - 10/2
}

func NewDerivedStats(stats BaseStats) DerivedStats {
	return DerivedStats{
		Mod: modifiers{
			Str: modifier{
				statify(stats.Str),
				P_NO_PROF,
			},
			Dex: modifier{
				statify(stats.Dex),
				P_NO_PROF,
			},
			Con: modifier{
				statify(stats.Con),
				P_NO_PROF,
			},
			Int: modifier{
				statify(stats.Int),
				P_NO_PROF,
			},
			Wis: modifier{
				statify(stats.Wis),
				P_NO_PROF,
			},
			Cha: modifier{
				statify(stats.Cha),
				P_NO_PROF,
			},
		},
		Init: 0,
		AC:   stats.Dex,
		SDC:  0,
		Skills: skills{
			Acrobatics:     modifier{},
			AnimalHandling: modifier{},
			Arcana:         modifier{},
			Athletics:      modifier{},
			Deception:      modifier{},
			History:        modifier{},
			Insight:        modifier{},
			Intimidation:   modifier{},
			Investigation:  modifier{},
			Medicine:       modifier{},
			Nature:         modifier{},
			Perception:     modifier{},
			Performance:    modifier{},
			Persuasion:     modifier{},
			Religion:       modifier{},
			SlightOfHand:   modifier{},
			Stealth:        modifier{},
			Survival:       modifier{},
		},
		Saves: saves{
			Str: modifier{
				statify(stats.Str),
				P_NO_PROF,
			},
			Dex: modifier{
				statify(stats.Dex),
				P_NO_PROF,
			},
			Con: modifier{
				statify(stats.Con),
				P_NO_PROF,
			},
			Int: modifier{
				statify(stats.Int),
				P_NO_PROF,
			},
			Wis: modifier{
				statify(stats.Wis),
				P_NO_PROF,
			},
			Cha: modifier{
				statify(stats.Cha),
				P_NO_PROF,
			},
		},
	}
}
