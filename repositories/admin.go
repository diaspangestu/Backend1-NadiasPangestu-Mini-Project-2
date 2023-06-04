package repositories

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	LoginAdmin(username string) (*entities.Actor, error)
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	RegisterAdmin(admin *entities.Actor) (*entities.Actor, error)
}

type Admin struct {
	db *gorm.DB
}

func NewAdmin(db *gorm.DB) Admin {
	return Admin{
		db: db,
	}
}

func (repo Admin) LoginAdmin(username string) (*entities.Actor, error) {
	admin := &entities.Actor{}

	err := repo.db.Model(&entities.Actor{}).
		Where("username = ? AND is_verified = ?", username, "true").First(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

// CreateCustomer Admin
func (repo Admin) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo Admin) RegisterAdmin(admin *entities.Actor) (*entities.Actor, error) {
	err := repo.db.Model(&entities.Actor{}).Create(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}