package main

import (
	"golang-graphql-talk/modules"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    modules.QueryType,
	Mutation: modules.MutationType,
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
	fs := http.FileServer(http.Dir("static"))

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	http.Handle("/", fs)

	// and serve!
	http.ListenAndServe(":"+port, nil)
}
