package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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

var linkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Link",
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

type link struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

func query(db shortcut.Database) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"link": &graphql.Field{
				Description: "Get link by key",
				Type:        linkType,
				Args:        graphql.FieldConfigArgument{"key": &graphql.ArgumentConfig{Type: graphql.String}},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					key, _ := params.Args["key"].(string)
					url, _ := db.Lookup(key)
					return link{Key: key, URL: url.String()}, nil
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

func NewGraphQL(database shortcut.Database) Schema {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: query(database),
	})

	return &GraphQL{Schema: schema}
}

func NewAPI(db shortcut.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := NewGraphQL(db).Execute(r.URL.Query().Get("query"))

		json.NewEncoder(w).Encode(result)
	})
}
