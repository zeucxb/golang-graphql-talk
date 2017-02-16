package types

import "github.com/graphql-go/graphql"

var PostQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PostQuery",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})
