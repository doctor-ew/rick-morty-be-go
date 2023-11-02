package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/doctor-ew/rick-morty-be-go/graphql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Create a new resolver instance
	resolver := graphql.Resolver{} // Replace with the actual resolver initialization

	// Create the GraphQL server and wire it up
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver}))

	r.Handle("/query", srv)
	r.Handle("/playground", playground.Handler("GraphQL Playground", "/query"))

	http.Handle("/", r)

	// Define the port to listen on (e.g., 4000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // Default port
	}

	http.ListenAndServe(":"+port, nil)
}
