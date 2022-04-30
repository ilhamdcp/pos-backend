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
	Name: "Vendor",
	Values: graphql.EnumValueConfigMap{
		"GRAB": &graphql.EnumValueConfig{
			Value: GRAB,
		},
		"GOJEK": &graphql.EnumValueConfig{
			Value: GOJEK,
		},
		"SHOPEE_FOOD": &graphql.EnumValueConfig{
			Value: SHOPEE_FOOD,
		},
		"UNKNOWN": &graphql.EnumValueConfig{
			Value: UNKNOWN,
		},
	},
})
