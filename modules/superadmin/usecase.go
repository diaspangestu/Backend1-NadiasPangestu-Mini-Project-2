package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/helpers"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"time"
)

type UsecaseSuperadminInterface interface {
	CreateSuperadmin(superadmin SuperAdminParam) (*entities.Actor, error)
	LoginSuperadmin(id uint, username, password string) (*entities.Actor, string, error)
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	DeleteCustomerById(id uint) error
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error)
	ApprovedAdminRegister(id uint) error
	RejectedAdminRegister(id uint) error
	UpdateActivedAdmin(id uint) error
	UpdateDeadactivedAdmin(id uint) error
	GetApprovalRequest() ([]*entities.Actor, error)
	GetAllAdmins(username string, page, pageSize int) ([]*entities.Actor, error)
}

type UsecaseSuperadmin struct {
	superadminRepo repositories.Superadmin
}

func (uc UsecaseSuperadmin) CreateSuperadmin(superadmin SuperAdminParam) (*entities.Actor, error) {
	hasPass := helpers.HashPass(superadmin.Password)

	newSuperadmin := &entities.Actor{
		Username:   superadmin.Username,
		Password:   hasPass,
		RoleID:     1,
		IsVerified: entities.True,
		IsActived:  entities.True,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createSuperadmin, err := uc.superadminRepo.CreateSuperadmin(newSuperadmin)
	if err != nil {
		return nil, err
	}

	return createSuperadmin, nil
}

func (uc UsecaseSuperadmin) LoginSuperadmin(id uint, username, password string) (*entities.Actor, string, error) {
	superAdmin, err := uc.superadminRepo.LoginSuperadmin(username)
	if err != nil {
		return nil, "", err
	}

	// Verify hashed password
	comparePass := helpers.ComparePass([]byte(superAdmin.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	// Generate token JWT
	tokenString, err := helpers.GenerateToken(id, username)
	if err != nil {
		return nil, "", err
	}

	return superAdmin, tokenString, nil
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
	// Get Existing Data Data
	existingData, err := uc.superadminRepo.GetCustomerById(id)
	if err != nil {
		return err
	}

	return uc.superadminRepo.DeleteCustomerById(id, existingData)
}

func (uc UsecaseSuperadmin) GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error) {
	customers, err := uc.superadminRepo.GetAllCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		return nil, err
	}

	return customers, nil
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

func (uc UsecaseSuperadmin) UpdateActivedAdmin(id uint) error {
	err := uc.superadminRepo.UpdateActivedAdmin(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) UpdateDeadactivedAdmin(id uint) error {
	err := uc.superadminRepo.UpdateDeadactivedAdmin(id)
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

func (uc UsecaseSuperadmin) GetAllAdmins(username string, page, pageSize int) ([]*entities.Actor, error) {
	admins, err := uc.superadminRepo.GetAllAdmins(username, page, pageSize)
	if err != nil {
		return nil, err
	}

	return admins, nil
}
