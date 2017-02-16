package resolvers

import (
	"golang-graphql-talk/helpers"

	"github.com/graphql-go/graphql"
)

var PostsResolver = func(_ graphql.ResolveParams) (interface{}, error) {
	return helpers.PostList, nil
}
