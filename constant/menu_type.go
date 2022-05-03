package constant

import "github.com/graphql-go/graphql"

type MenuCategory int

const (
	MAIN_COURSE MenuCategory = iota
	SIDE_DISH                = iota
	COFFEE                   = iota
	NON_COFFEE               = iota
)

var MenuCategoryGql = graphql.NewEnum(graphql.EnumConfig{
	Name:        "MenuCategory",
	Description: "Category of the menu",
	Values: graphql.EnumValueConfigMap{
		"MainCourse": &graphql.EnumValueConfig{
			Value: MAIN_COURSE,
		},
		"SideDish": &graphql.EnumValueConfig{
			Value: SIDE_DISH,
		},
		"Coffee": &graphql.EnumValueConfig{
			Value: COFFEE,
		},
		"NonCofee": &graphql.EnumValueConfig{
			Value: NON_COFFEE,
		},
	},
})
