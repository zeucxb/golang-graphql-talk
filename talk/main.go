package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

/*
 *
 * Tipo para amazenar os posts
 *
 */
type Post struct {
	Message string `json:"message"`
	Rate    int    `json:"rate"`
}

var PostList []Post

/*
 *
 * Tipo que representa um post
 *
 */
var PostQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PostQuery",
	Description: "Objeto que representa uma postagem",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type:        graphql.String,
			Description: "Mensagem da postagem",
		},
		"rate": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

/*
 *
 * Resolve do field posts
 *
 */
var PostsResolver = func(_ graphql.ResolveParams) (interface{}, error) {
	return PostList, nil
}

/*
 *
 * Tipo Query, root da Query
 *
 */
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Query do graphQL",
	Fields: graphql.Fields{
		"posts": &graphql.Field{
			Type:        graphql.NewList(PostQueryType),
			Resolve:     PostsResolver,
			Description: "Todos os posts",
		},
	},
})

/*
 *
 * Resolve do mutation insertPost
 *
 */
var InsertPostResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newMessage, ok := params.Args["message"].(string)
	if !ok {
		return newMessage, fmt.Errorf("SYSTEM ERROR")
	}

	newRate, okRate := params.Args["rate"].(int)
	if !okRate {
		return newMessage, fmt.Errorf("SYSTEM ERROR")
	}

	newPost := Post{
		Message: newMessage,
		Rate:    newRate,
	}

	PostList = append([]Post{newPost}, PostList...)

	return newPost.Message, nil
}

/*
 *
 * Tipo Mutation, root da Mutation
 *
 */
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Mutation",
	Description: "Mutation do graphQL",
	Fields: graphql.Fields{
		"insertPost": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"message": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"rate": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve:     InsertPostResolver,
			Description: "Adiciona um novo post",
		},
	},
})

/*
 *
 * Schema
 *
 */
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    QueryType,
	Mutation: MutationType,
})

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// create a graphql-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("../static"))

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	http.Handle("/", fs)

	// and serve!
	http.ListenAndServe(":"+port, nil)
}
