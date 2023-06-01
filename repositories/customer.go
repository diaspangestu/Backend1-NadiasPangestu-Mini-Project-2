package repositories

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	GetCustomerById(id uint) (*entities.Customer, error)
	UpdateCustomerById(id uint, customer *entities.Customer) (*entities.Customer, error)
	DeleteCustomerById(id uint, customer *entities.Customer) error
}

type Customer struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) Customer {
	return Customer{
		db: db,
	}
}

func (repo Customer) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo Customer) GetCustomerById(id uint) (*entities.Customer, error) {
	customer := &entities.Customer{}

	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).First(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo Customer) UpdateCustomerById(id uint, customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).Save(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo Customer) DeleteCustomerById(id uint, customer *entities.Customer) error {
	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).Delete(customer).Error
	if err != nil {
		return err
	}

	return nil
}
