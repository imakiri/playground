package api

import (
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var foo = graphql.Fields{
	"doo": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "Hello World!", nil
		},
		Description: "",
	},
	"moo": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "Hello!", nil
		},
	},
}

var rQuery = graphql.ObjectConfig{
	Name: "root",
	Fields: graphql.Fields{
		"foo": &graphql.Field{
			Type:              foo,
			Args:              nil,
			Resolve:           nil,
			DeprecationReason: "",
			Description:       "",
		},
	},
}

var schemaCfg = graphql.SchemaConfig{
	Query: graphql.NewObject(rQuery),
}

var schema, _ = graphql.NewSchema(schemaCfg)

func RunGraphQL(rr *mux.Router) error {
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})

	rr.Handle("/graphql", h)

	return nil
}
