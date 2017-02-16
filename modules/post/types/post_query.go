package types

import "github.com/graphql-go/graphql"

var PostQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PostQuery",
	Description: "Objeto que representa uma postagem",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type:        graphql.String,
			Description: "Mensagem da postagem",
		},
	},
})
