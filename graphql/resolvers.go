package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

type Character struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Species  string    `json:"species"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Image    string    `json:"image"`
	Episodes []Episode `json:"episode"`
}

type Episode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RickAndMortyAssociation struct {
	Rick    Character   `json:"rick"`
	Morties []Character `json:"morties"`
}

var (
	// Define your query type here
	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"charactersByName": &graphql.Field{
				Type:        graphql.NewList(characterType),
				Description: "Get characters by name",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Implement your resolver logic here
					// You can use the Redis client to cache data
					name, _ := p.Args["name"].(string)
					fmt.Printf("Fetching characters by name: %s\n", name)
					// Implement the character retrieval logic
					// Return characters as []Character
					return []Character{}, nil
				},
			},
			// Define other queries here
		},
	})
)
