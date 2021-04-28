package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"com.damonkelley/shortcut/internal/shortcut"
	"github.com/graphql-go/graphql"
)

type Result struct {
	*graphql.Result
}

type Schema interface {
	Execute(query string) Result
}

type GraphQL struct {
	graphql.Schema
}

var shortcutType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "shortcut",
		Fields: graphql.Fields{
			"key": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

type externalShortcut struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func mutate(db shortcut.Put) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutate",
		Fields: graphql.Fields{
			"generate": &graphql.Field{
				Description: "Generate a shortcut",
				Type:        shortcutType,
				Args:        graphql.FieldConfigArgument{"url": &graphql.ArgumentConfig{Type: graphql.String}},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rawURl, _ := params.Args["url"].(string)
					url, err := url.Parse(rawURl)

					if err != nil {
						return externalShortcut{}, err
					}

					key := db.Put(url)
					return externalShortcut{Key: key, URL: url.String()}, nil
				},
			},
		},
	})
}

func query(db shortcut.Lookup) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"shortcut": &graphql.Field{
				Description: "Get shortcut by key",
				Type:        shortcutType,
				Args:        graphql.FieldConfigArgument{"key": &graphql.ArgumentConfig{Type: graphql.String}},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					key, _ := params.Args["key"].(string)
					url, err := db.Lookup(key)
					if err != nil {
						return externalShortcut{}, err
					}
					return externalShortcut{Key: key, URL: url.String()}, nil
				},
			},
		},
	})
}

func (schema *GraphQL) Execute(query string) Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}

	return Result{result}

}

func NewGraphQL(database shortcut.ReadWrite) Schema {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query(database),
		Mutation: mutate(database),
	})

	return &GraphQL{Schema: schema}
}

func NewAPI(db shortcut.ReadWrite) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := NewGraphQL(db).Execute(r.URL.Query().Get("query"))

		json.NewEncoder(w).Encode(result)
	})
}
