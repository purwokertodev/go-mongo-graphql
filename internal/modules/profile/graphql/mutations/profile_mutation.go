package mutations

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/graphql/types"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/model"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/repository"
)

//ProfileMutation
type ProfileMutation struct {
	profileRepository repository.ProfileRepository
}

//New profileMutation's Constructor
func New(profileRepository repository.ProfileRepository) *ProfileMutation {
	return &ProfileMutation{profileRepository}
}

//GetMutation return graphql object
func (pm *ProfileMutation) GetMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createProfile": &graphql.Field{
				Type:        types.ProfileType,
				Description: "Create New Profile",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"firstName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"lastName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, _ := params.Args["id"].(string)
					firstName, _ := params.Args["firstName"].(string)
					lastName, _ := params.Args["lastName"].(string)
					email, _ := params.Args["email"].(string)
					password, _ := params.Args["password"].(string)

					var p model.Profile

					p.ID = id
					p.FirstName = firstName
					p.LastName = lastName
					p.Email = email
					p.Password = password
					p.CreatedAt = time.Now()
					p.UpdatedAt = time.Now()

					err := pm.profileRepository.Save(&p)

					if err != nil {
						return nil, err
					}

					return p, nil
				},
			},
		},
	})
}
