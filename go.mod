module github.com/kjkondratuk/goblins-and-gold

go 1.18

require (
	github.com/google/uuid v1.3.0
	github.com/olekukonko/ts v0.0.0-20171002115256-78ecb04241c0
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/atomicgo/cursor v0.0.1 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/gookit/color v1.4.2 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/pterm/pterm v0.12.39 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	golang.org/x/sys v0.0.0-20220325203850-36772127a21f // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/goccy/go-yaml v1.9.5
	github.com/pmezard/go-difflib v1.0.0 // indirect
)

// Use my custom version of the lib that has
replace github.com/goccy/go-yaml v1.9.5 => github.com/kjkondratuk/go-yaml v1.9.6
