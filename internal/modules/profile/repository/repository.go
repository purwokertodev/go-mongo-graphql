package repository

import (
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/model"
)

//ProfileRepository interface
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindByEmail(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
