package superadmin

import (
	"errors"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"time"
)

type UsecaseSuperadminInterface interface {
	LoginSuperadmin(username, password string) (*entities.Actor, error)
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	DeleteCustomerById(id uint) error
	ApprovedAdminRegister(id uint) error
	RejectedAdminRegister(id uint) error
	GetApprovalRequest() ([]*entities.Actor, error)
}

type UsecaseSuperadmin struct {
	superadminRepo repositories.Superadmin
}

func (uc UsecaseSuperadmin) LoginSuperadmin(username, password string) (*entities.Actor, error) {
	superAdmin, err := uc.superadminRepo.LoginSuperadmin(username)
	if err != nil {
		return nil, err
	}

	// Check password
	if superAdmin.Password != password {
		return nil, errors.New("wrong password")
	}

	return superAdmin, nil
}

// CreateCustomer Superadmin
func (uc UsecaseSuperadmin) CreateCustomer(customer CustomerParam) (entities.Customer, error) {
	var newCustomer *entities.Customer

	newCustomer = &entities.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.superadminRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}

	return *newCustomer, nil
}

// DeleteCustomerById Superadmin
func (uc UsecaseSuperadmin) DeleteCustomerById(id uint) error {
	// Get Existing Customer Data
	existingData, err := uc.superadminRepo.GetCustomerById(id)
	if err != nil {
		return err
	}

	return uc.superadminRepo.DeleteCustomerById(id, existingData)
}

func (uc UsecaseSuperadmin) ApprovedAdminRegister(id uint) error {
	err := uc.superadminRepo.ApprovedAdminRegister(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) RejectedAdminRegister(id uint) error {
	err := uc.superadminRepo.RejectedAdminRegister(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) GetApprovalRequest() ([]*entities.Actor, error) {
	request, err := uc.superadminRepo.GetApprovalRequest()
	if err != nil {
		return nil, err
	}

	return request, nil
}
