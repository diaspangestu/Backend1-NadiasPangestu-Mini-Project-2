package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"time"
)

type UsecaseAdminInterface interface {
	RegisterAdmin(admin AdminParam) (entities.Actor, error)
}

type UsecaseAdmin struct {
	adminRepo repositories.Admin
}

func (uc UsecaseAdmin) RegisterAdmin(admin AdminParam) (entities.Actor, error) {
	var newAdmin *entities.Actor

	newAdmin = &entities.Actor{
		Username:   admin.Username,
		Password:   admin.Password,
		RoleID:     2,
		IsVerified: entities.False,
		IsActived:  entities.False,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err := uc.adminRepo.RegisterAdmin(newAdmin)
	if err != nil {
		return *newAdmin, err
	}

	return *newAdmin, nil
}
