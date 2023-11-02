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

// Define GraphQL object type for Character
var characterType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Character",
	Fields: graphql.Fields{
		"ID":       &graphql.Field{Type: graphql.String},
		"Name":     &graphql.Field{Type: graphql.String},
		"Status":   &graphql.Field{Type: graphql.String},
		"Species":  &graphql.Field{Type: graphql.String},
		"Type":     &graphql.Field{Type: graphql.String},
		"Gender":   &graphql.Field{Type: graphql.String},
		"Image":    &graphql.Field{Type: graphql.String},
		"Episodes": &graphql.Field{Type: graphql.NewList(episodeType)}, // Assuming you'll define episodeType similarly
	},
})

// Define GraphQL object type for Episode (assuming you'll use it later)
var episodeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Episode",
	Fields: graphql.Fields{
		"ID":   &graphql.Field{Type: graphql.String},
		"Name": &graphql.Field{Type: graphql.String},
	},
})

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
					name, ok := p.Args["name"].(string)
					if !ok {
						return nil, fmt.Errorf("name argument is missing or not a string")
					}
					fmt.Printf("Fetching characters by name: %s\n", name)
					// Implement the character retrieval logic
					// For now, returning an empty list
					// In a real-world scenario, you'd fetch this from a database or API
					return []Character{}, nil
				},
			},
			// Define other queries here
		},
	})
)
