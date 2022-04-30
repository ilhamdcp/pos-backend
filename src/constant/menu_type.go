package constant

import "github.com/graphql-go/graphql"

type MenuType int

const (
	MAIN_COURSE MenuType = iota
	SIDE_DISH            = iota
	COFFEE               = iota
	NON_COFFEE           = iota
)

var MenuTypeGql = graphql.NewEnum(graphql.EnumConfig{
	Name: "Menu Type",
	Values: graphql.EnumValueConfigMap{
		"Main` Course": &graphql.EnumValueConfig{
			Value: MAIN_COURSE,
		},
		"Side Dish": &graphql.EnumValueConfig{
			Value: SIDE_DISH,
		},
		"Coffee": &graphql.EnumValueConfig{
			Value: COFFEE,
		},
		"Non Coffee": &graphql.EnumValueConfig{
			Value: NON_COFFEE,
		},
	},
})
