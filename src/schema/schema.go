package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"pos-backend/constant"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type Customer struct {
	Id     int             `json:"id"`
	Name   string          `json:"name"`
	Driver string          `json:"driver"`
	Vendor constant.Vendor `json:"vendor"`
}

type Order struct {
	Id            string      `json:"id"`
	Customer      Customer    `json:"customer"`
	TotalQuantity int         `json:"total_quantity"`
	TotalPrice    float32     `json:"total_price"`
	Item          []OrderItem `json:"items"`
	IsVat         bool        `json:"is_vat"`
	Vat           float32     `json:"vat"`
}

type OrderItem struct {
	Id            string  `json:"_id"`
	OrderId       int     `json:"order_id"`
	Menu          Menu    `json:"menu"`
	Quantity      int     `json:"quantity"`
	SubTotalPrice float32 `json:"subtotal_price"`
}

type Menu struct {
	Id          string  `json:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Type        string  `json:"type"`
}

//var userType = graphql.NewObject(graphql.ObjectConfig{
//	Name: "User",
//	Fields: graphql.Fields{
//		"username": &graphql.Field{
//			Type: graphql.String,
//		},
//		"name": &graphql.Field{
//			Type: graphql.String,
//		},
//	},
//})
//
//var customerType = graphql.NewObject(graphql.ObjectConfig{
//	Name: "Customer",
//	Fields: graphql.Fields{
//		"id": &graphql.Field{
//			Type: graphql.Int,
//		},
//		"name": &graphql.Field{
//			Type: graphql.String,
//		},
//		"driver": &graphql.Field{
//			Type: graphql.String,
//		},
//		"vendor": &graphql.Field{
//			Type: constant.VendorType,
//		},
//	},
//})

var menuType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Menu",
	Fields: &graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"type": &graphql.Field{
			Type: constant.MenuTypeGql,
		},
	},
})

//var orderType = graphql.NewObject(graphql.ObjectConfig{
//	Name: "Order",
//	Fields: graphql.Fields{
//		"id": &graphql.Field{
//			Type: graphql.Int,
//		},
//		"customer": &graphql.Field{
//			Type: customerType,
//		},
//	},
//})

var queries = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getMenus": &graphql.Field{
			Type: graphql.NewList(menuType),
			Args: graphql.FieldConfigArgument{
				"category": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Category of the menu",
				},
				"keyword": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "Keyword that matches either the menu name or description",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
		"getMenuById": &graphql.Field{
			Type: menuType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "ID of the menu",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, errors.New("Unimplemented")
			},
		},
	},
})

//var mutations = graphql.NewObject(graphql.ObjectConfig{})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries,
})
