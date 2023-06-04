package repositories

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"gorm.io/gorm"
)

type SuperAdminRepository interface {
	LoginSuperadmin(username string) (*entities.Actor, error)
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	GetCustomerById(id uint) (*entities.Customer, error)
	DeleteCustomerById(id uint, customer *entities.Customer) error
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error)
	ApprovedAdminRegister(id uint) error
	RejectedAdminRegister(id uint) error
	UpdateActivedAdmin(id uint) error
	UpdateDeadactivedAdmin(id uint) error
	GetApprovalRequest() ([]*entities.Actor, error)
}

type Superadmin struct {
	db *gorm.DB
}

func NewSuperadmin(db *gorm.DB) Superadmin {
	return Superadmin{
		db: db,
	}
}

func (repo Superadmin) LoginSuperadmin(username string) (*entities.Actor, error) {
	superAdmin := &entities.Actor{}

	err := repo.db.Model(&entities.Actor{}).Where("username = ?", username).First(superAdmin).Error
	if err != nil {
		return nil, err
	}

	return superAdmin, nil
}

// CreateCustomer Superadmin can create customer data
func (repo Superadmin) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := repo.db.Model(&entities.Customer{}).Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// GetCustomerById Superadmin
func (repo Superadmin) GetCustomerById(id uint) (*entities.Customer, error) {
	customer := &entities.Customer{}

	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).First(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// DeleteCustomerById Superadmin
func (repo Superadmin) DeleteCustomerById(id uint, customer *entities.Customer) error {
	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).Delete(customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo Superadmin) GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error) {
	var customers []*entities.Customer

	query := repo.db.Model(&entities.Customer{})
	if first_name != "" {
		query = query.Where("first_name LIKE ?", "%"+first_name+"%")
	} else if last_name != "" {
		query = query.Where("last_name LIKE ?", "%"+last_name+"%")
	} else if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	// Pagination
	offset := (page - 1) * pageSize

	err := query.Offset(offset).Limit(pageSize).Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repo Superadmin) ApprovedAdminRegister(id uint) error {
	// Check admin data by id
	err := repo.db.Where("id = ?", id).First(&entities.Actor{}).Error
	if err != nil {
		return err
	}

	// Update Verified
	err = repo.db.Model(&entities.Actor{}).Where("id = ?", id).Update("is_verified", true).Error
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

func (repo Superadmin) UpdateActivedAdmin(id uint) error {
	// Check admin data by id
	err := repo.db.Where("id = ?", id).First(&entities.Actor{}).Error
	if err != nil {
		return err
	}

	admin := &entities.Actor{}

	// Update Actived
	if admin.IsActived != "true" {
		err = repo.db.Model(&entities.Actor{}).Where("id = ?", id).Update("is_actived", "true").Error
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (repo Superadmin) UpdateDeadactivedAdmin(id uint) error {
	// Check admin data by id
	err := repo.db.Where("id = ?", id).First(&entities.Actor{}).Error
	if err != nil {
		return err
	}

	admin := &entities.Actor{}

	// Update Actived
	if admin.IsActived != "false" {
		err = repo.db.Model(&entities.Actor{}).Where("id = ?", id).Update("is_actived", "false").Error
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (repo Superadmin) GetApprovalRequest() ([]*entities.Actor, error) {
	var result []*entities.Actor

	err := repo.db.Model(&entities.Actor{}).
		Select("id, username, role_id, is_verified, created_at, updated_at").
		Where("role_id = ? AND is_verified = ?", 2, "false").
		Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
