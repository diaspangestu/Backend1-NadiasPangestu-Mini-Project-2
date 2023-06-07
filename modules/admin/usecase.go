package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/helpers"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"time"
)

type UsecaseAdminInterface interface {
	// admin
	LoginAdmin(username, password string) (entities.Actor, error)
	RegisterAdmin(admin LoginAdminParam) (entities.Actor, error)
	GetAdminById(id uint) (entities.Actor, error)
	UpdateCustomerById(id uint, admin AdminParam) (entities.Actor, error)
	DeleteAdminById(id uint) error
	// customer
	DeleteCustomerById(id uint) error
	CreateCustomer(customer CustomerParam) (entities.Customer, string, error)
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error)
	SaveCustomersFromAPI() error
}

type UsecaseAdmin struct {
	adminRepo repositories.Admin
}

func (uc UsecaseAdmin) LoginAdmin(id uint, username, password string) (*entities.Actor, string, error) {
	admin, err := uc.adminRepo.LoginAdmin(username)
	if err != nil {
		return nil, "", err
	}

	// Verify hashed password
	comparePass := helpers.ComparePass([]byte(admin.Password), []byte(password))
	if !comparePass {
		return nil, "", err
	}

	// Generate token JWT
	tokenString, err := helpers.GenerateToken(id, username)
	if err != nil {
		return nil, "", err
	}

	return admin, tokenString, nil
}

func (uc UsecaseAdmin) RegisterAdmin(admin LoginAdminParam) (*entities.Actor, error) {
	hashPass := helpers.HashPass(admin.Password)

	newAdmin := &entities.Actor{
		Username:   admin.Username,
		Password:   hashPass,
		RoleID:     2,
		IsVerified: entities.False,
		IsActived:  entities.False,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createAdmin, err := uc.adminRepo.RegisterAdmin(newAdmin)
	if err != nil {
		return nil, err
	}

	return createAdmin, nil
}

func (uc UsecaseAdmin) GetAdminById(id uint) (entities.Actor, error) {
	admin, err := uc.adminRepo.GetAdminById(id)
	if err != nil {
		return entities.Actor{}, err
	}

	return *admin, nil
}

func (uc UsecaseAdmin) UpdateCustomerById(id uint, admin AdminParam) (entities.Actor, error) {
	hashPass := helpers.HashPass(admin.Password)

	// Get Existing Admin Data
	existingAdmin, err := uc.adminRepo.GetAdminById(id)
	if err != nil {
		return entities.Actor{}, err
	}

	existingAdmin.Username = admin.Username
	existingAdmin.Password = hashPass
	existingAdmin.UpdatedAt = time.Now()

	// Updated the Admin Data
	updatedAdmin, err := uc.adminRepo.UpdateAdminById(id, existingAdmin)
	if err != nil {
		return entities.Actor{}, err
	}

	return *updatedAdmin, nil
}

func (uc UsecaseAdmin) DeleteAdminById(id uint) error {
	// Get existing Admin Data
	existingData, err := uc.adminRepo.GetAdminById(id)
	if err != nil {
		return err
	}

	return uc.adminRepo.DeleteAdminById(id, existingData)
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

// DeleteCustomerById Admin
func (uc UsecaseAdmin) DeleteCustomerById(id uint) error {
	// Get Existing Customer Data
	existingData, err := uc.adminRepo.GetCustomerById(id)
	if err != nil {
		return err
	}

	return uc.adminRepo.DeleteCustomerById(id, existingData)
}

func (uc UsecaseAdmin) GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error) {
	customers, err := uc.adminRepo.GetAllCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (uc UsecaseAdmin) SaveCustomersFromAPI() error {
	url := "https://reqres.in/api/users?page=2"

	err := uc.adminRepo.SaveCustomersFromAPI(url)
	if err != nil {
		return err
	}

	return nil
}
