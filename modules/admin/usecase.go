package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"time"
)

type UsecaseAdminInterface interface {
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	RegisterAdmin(admin AdminParam) (entities.Actor, error)
}

type UsecaseAdmin struct {
	adminRepo repositories.Admin
}

// CreateCustomer Admin
func (uc UsecaseAdmin) CreateCustomer(customer CustomerParam) (entities.Customer, error) {
	var newCustomer *entities.Customer

	newCustomer = &entities.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.adminRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}

	return *newCustomer, nil
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
