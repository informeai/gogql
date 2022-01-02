package controllers

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/informeai/gogql/schemas"
)

type GraphQlController struct {
	schema graphql.Schema
}

func NewGraphQlController() *GraphQlController {
	return &GraphQlController{}
}
func (g *GraphQlController) setup() error {
	schemaConfig := graphql.SchemaConfig{Query: schemas.NewUserSchema().Query(), Mutation: schemas.NewUserSchema().Mutation()}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return err
	}
	g.schema = schema
	return nil
}

func (g *GraphQlController) Exec(query string) (string, error) {
	if err := g.setup(); err != nil {
		return "", err
	}
	params := graphql.Params{Schema: g.schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		return "", r.Errors[0]
	}
	result, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
