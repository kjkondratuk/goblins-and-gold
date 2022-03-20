package stats

const (
	defaultStat = 10
)

type baseStats struct {
	_lvl int
	_str int
	_dex int
	_con int
	_int int
	_wis int
	_cha int
}

type BaseStats interface {
	Lvl() int
	LevelUp() int
	Str() int
	Dex() int
	Con() int
	Int() int
	Wis() int
	Cha() int
}

type Option func(stats baseStats) baseStats

func NewBaseStats(opts ...Option) BaseStats {
	s := baseStats{
		_lvl: 1,
		_str: defaultStat,
		_dex: defaultStat,
		_con: defaultStat,
		_int: defaultStat,
		_wis: defaultStat,
		_cha: defaultStat,
	}

	for _, opt := range opts {
		s = opt(s)
	}

	return &s
}

func WithLvl(lvl int) Option {
	return func(stats baseStats) baseStats {
		stats._lvl = lvl
		return stats
	}
}

func WithStr(str int) Option {
	return func(stats baseStats) baseStats {
		stats._str = str
		return stats
	}
}

func WithDex(dex int) Option {
	return func(stats baseStats) baseStats {
		stats._dex = dex
		return stats
	}
}

func WithCon(con int) Option {
	return func(stats baseStats) baseStats {
		stats._con = con
		return stats
	}
}

func WithInt(i int) Option {
	return func(stats baseStats) baseStats {
		stats._int = i
		return stats
	}
}

func WithWis(wis int) Option {
	return func(stats baseStats) baseStats {
		stats._wis = wis
		return stats
	}
}

func WithCha(cha int) Option {
	return func(stats baseStats) baseStats {
		stats._cha = cha
		return stats
	}
}

func (s *baseStats) Lvl() int {
	return s._lvl
}

func (s *baseStats) LevelUp() int {
	s._lvl += 1
	return s._lvl
}

func (s *baseStats) Str() int {
	return s._str
}

func (s *baseStats) Dex() int {
	return s._dex
}

func (s *baseStats) Con() int {
	return s._con
}

func (s *baseStats) Int() int {
	return s._int
}

func (s *baseStats) Wis() int {
	return s._wis
}

func (s *baseStats) Cha() int {
	return s._cha
}
