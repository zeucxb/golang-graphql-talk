package resolvers

import (
	"fmt"
	"golang-graphql-talk/helpers"

	"github.com/graphql-go/graphql"
)

var InsertPostResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newMessage, ok := params.Args["message"].(string)
	if !ok {
		return newMessage, fmt.Errorf("SYSTEM ERROR")
	}

	newPost := helpers.Post{
		Message: newMessage,
	}

	helpers.PostList = append([]helpers.Post{newPost}, helpers.PostList...)

	return newPost.Message, nil
}
