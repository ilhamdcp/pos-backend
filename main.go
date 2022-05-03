package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
	"pos-backend/schema"
)

func main() {
	//http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	//	query := r.URL.Query().Get("query")
	//	result := graphql.Do(graphql.Params{Schema: schema.Schema, RequestString: query})
	//	err := json.NewEncoder(w).Encode(result)
	//	if err != nil {
	//		return
	//	}
	//})
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	return
	//}
	//fmt.Println("Graphql API is listening on port 8080")
	var schema, error = graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Queries,
	})
	if error != nil {
		fmt.Println(error)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
