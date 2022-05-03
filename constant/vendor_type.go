package constant

import "github.com/graphql-go/graphql"

type Vendor int

const (
	GRAB        Vendor = iota
	GOJEK              = iota
	SHOPEE_FOOD        = iota
	UNKNOWN            = iota
)

var VendorType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Vendor",
	Description: "The vendor which facilitate the order",
	Values: graphql.EnumValueConfigMap{
		"Grab": &graphql.EnumValueConfig{
			Value: GRAB,
		},
		"Gojek": &graphql.EnumValueConfig{
			Value: GOJEK,
		},
		"ShopeeFood": &graphql.EnumValueConfig{
			Value: SHOPEE_FOOD,
		},
		"Unknown": &graphql.EnumValueConfig{
			Value: UNKNOWN,
		},
	},
})
