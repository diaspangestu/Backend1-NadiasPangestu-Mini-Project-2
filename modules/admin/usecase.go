package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UsecaseAdminInterface interface {
	LoginAdmin(username, password string) (entities.Actor, error)
	RegisterAdmin(admin AdminParam) (entities.Actor, error)
	DeleteCustomerById(id uint) error
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error)
}

type UsecaseAdmin struct {
	adminRepo repositories.Admin
}

func (uc UsecaseAdmin) LoginAdmin(username, password string) (*entities.Actor, error) {
	admin, err := uc.adminRepo.LoginAdmin(username)
	if err != nil {
		return nil, err
	}

	// Verify hashed password
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (uc UsecaseAdmin) RegisterAdmin(admin AdminParam) (*entities.Actor, error) {
	salt := 16
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), salt)
	if err != nil {
		return nil, err
	}

	newAdmin := &entities.Actor{
		Username:   admin.Username,
		Password:   string(hashedPassword),
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
