package user

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
)

type Usecase struct {
	userRepo repositories.UserRepositoryInterface
}

type UsecaseInterface interface {
	GetUserByID(payload Payload) []entities.User
}

func (uc Usecase) GetUserByID(payload Payload) []entities.User {
	user := uc.userRepo.GetByID(payload.ID)

	// if len user == 0 return no user
	if len(user) == 0 {
		return nil
	}

	return user
}
