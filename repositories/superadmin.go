package repositories

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"gorm.io/gorm"
)

type SuperAdminRepository interface {
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	ApprovedAdminRegister(id uint) error
	RejectedAdminRegister(id uint) error
}

type Superadmin struct {
	db *gorm.DB
}

func NewSuperadmin(db *gorm.DB) Superadmin {
	return Superadmin{
		db: db,
	}
}

// CreateCustomer Superadmin can create customer data
func (repo Superadmin) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repo Superadmin) ApprovedAdminRegister(id uint) error {
	// Check admin data by id
	err := repo.db.Where("id = ?", id).First(&entities.Actor{}).Error
	if err != nil {
		return err
	}

	// Update Verified and Actived
	err = repo.db.Model(&entities.Actor{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_verified": true,
		"is_actived":  true,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo Superadmin) RejectedAdminRegister(id uint) error {
	// Check admin data by id
	err := repo.db.Where("id = ?", id).First(&entities.Actor{}).Error
	if err != nil {
		return err
	}

	return nil
}
