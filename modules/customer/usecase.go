package customer

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
	"time"
)

type UsecaseCustomerInterface interface {
	CreateCustomer(customer CustomerParam) (entities.Customer, error)
	GetCustomerById(id uint) (entities.Customer, error)
	UpdateCustomerById(id uint, customer CustomerParam) (entities.Customer, error)
	DeleteCustomerById(id uint) error
}

type UsecaseCustomer struct {
	customerRepo repositories.CustomerRepositoryInterface
}

func (uc UsecaseCustomer) CreateCustomer(customer CustomerParam) (entities.Customer, error) {
	var newCustomer *entities.Customer

	newCustomer = &entities.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}

func (uc UsecaseCustomer) GetCustomerById(id uint) (entities.Customer, error) {
	customer, err := uc.customerRepo.GetCustomerById(id)
	if err != nil {
		return entities.Customer{}, err
	}

	return *customer, nil
}

func (uc UsecaseCustomer) UpdateCustomerById(id uint, customer CustomerParam) (entities.Customer, error) {
	// Get Existing Customer Data
	existingCustomer, err := uc.customerRepo.GetCustomerById(id)
	if err != nil {
		return entities.Customer{}, err
	}

	existingCustomer.FirstName = customer.FirstName
	existingCustomer.LastName = customer.LastName
	existingCustomer.Email = customer.Email
	existingCustomer.Avatar = customer.Avatar
	existingCustomer.UpdatedAt = time.Now()

	// Updated the Customer Data
	updatedCustomer, err := uc.customerRepo.UpdateCustomerById(id, existingCustomer)
	if err != nil {
		return entities.Customer{}, err
	}

	return *updatedCustomer, nil
}

func (uc UsecaseCustomer) DeleteCustomerById(id uint) error {
	// Get existing Customer Data
	existingData, err := uc.customerRepo.GetCustomerById(id)
	if err != nil {
		return err
	}

	return uc.customerRepo.DeleteCustomerById(id, existingData)
}
