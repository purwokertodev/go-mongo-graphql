package queries

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/graphql/types"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/repository"
)

//ProfileQuery
type ProfileQuery struct {
	profileRepository repository.ProfileRepository
}

//New profileQuery's Constructor
func New(profileRepository repository.ProfileRepository) *ProfileQuery {
	return &ProfileQuery{profileRepository}
}

//GetQuery return graphql object
func (pm *ProfileQuery) GetQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"profile": &graphql.Field{
				Type:        types.ProfileType,
				Description: "Get Profile",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, ok := params.Args["id"].(string)
					if !ok {
						return nil, errors.New("Require ID param")
					}

					profile, err := pm.profileRepository.FindByID(id)

					if err != nil {
						return nil, err
					}

					return profile, nil

				},
			},
			"profiles": &graphql.Field{
				Type:        graphql.NewList(types.ProfileType),
				Description: "Get Profiles",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					profiles, err := pm.profileRepository.FindAll()

					if err != nil {
						return nil, err
					}

					return profiles, nil
				},
			},
		},
	})
}
