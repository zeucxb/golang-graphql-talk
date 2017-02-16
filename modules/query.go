package modules

import (
	"golang-graphql-talk/modules/post/resolvers"
	"golang-graphql-talk/modules/post/types"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"posts": &graphql.Field{
			Type:    graphql.NewList(types.PostQueryType),
			Resolve: resolvers.PostsResolver,
		},
	},
})
