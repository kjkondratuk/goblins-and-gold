package ux

type Described interface {
	Describe() string
}

func DescribeToList(items []Described) []string {
	list := make([]string, 0)
	for _, i := range items {
		list = append(list, i.Describe())
	}
	return list
}
