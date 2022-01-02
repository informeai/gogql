package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/informeai/gogql/mocks"
)

var usersMock = []mocks.UserMock{
	{Id: 1, Name: "wellington", Email: "contato@gmail.com"},
}
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

type UserSchema struct{}

func NewUserSchema() *UserSchema {
	return &UserSchema{}
}

func (u *UserSchema) Query() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userType,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//use repository pattern
					id, ok := p.Args["id"].(int)
					if ok {
						for _, u := range usersMock {
							if u.Id == id {
								return u, nil
							}
						}
					}
					return nil, nil
				},
			},
			"users": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "Get all users",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return usersMock, nil
				},
			},
		},
	})
}

func (u *UserSchema) Mutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        userType,
				Description: "create user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//use repository pattern
					user := mocks.UserMock{
						Id:    p.Args["id"].(int),
						Name:  p.Args["name"].(string),
						Email: p.Args["email"].(string),
					}
					usersMock = append(usersMock, user)
					return user, nil
				},
			},
		},
	})
}
