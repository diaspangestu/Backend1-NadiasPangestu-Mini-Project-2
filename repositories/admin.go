package repositories

import (
	"encoding/json"
	"errors"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type AdminRepositoryInterface interface {
	LoginAdmin(username string) (*entities.Actor, error)
	RegisterAdmin(admin *entities.Actor) (*entities.Actor, error)
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	GetCustomerById(id uint) (*entities.Customer, error)
	GetCustomerByEmail(email string) (*entities.Customer, error)
	DeleteCustomerById(id uint, customer *entities.Customer) error
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error)
	FetchCustomersFromAPI() ([]*entities.Customer, error)
	SaveCustomersFromAPI(url string) error
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

	err := repo.db.Model(&entities.Actor{}).Where("username = ? AND is_verified = ? AND is_actived = ?", username, "true", "true").First(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (repo Admin) RegisterAdmin(admin *entities.Actor) (*entities.Actor, error) {
	err := repo.db.Model(&entities.Actor{}).Create(admin).Error
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

// GetCustomerById Admin
func (repo Admin) GetCustomerById(id uint) (*entities.Customer, error) {
	customer := &entities.Customer{}

	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).First(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// DeleteCustomerById Superadmin
func (repo Admin) DeleteCustomerById(id uint, customer *entities.Customer) error {
	err := repo.db.Model(&entities.Customer{}).Where("id = ?", id).Delete(customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo Admin) GetAllCustomers(first_name, last_name, email string, page, pageSize int) ([]*entities.Customer, error) {
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

func (repo Admin) GetCustomerByEmail(email string) (*entities.Customer, error) {
	customer := &entities.Customer{}

	err := repo.db.Model(&entities.Customer{}).Where("email = ?", email).First(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

type Get struct {
	Data []entities.Customer `json:"data"`
}

func (repo Admin) SaveCustomersFromAPI(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	customersAPIResponse := new(Get)

	err = json.Unmarshal(body, customersAPIResponse)
	if err != nil {
		return err
	}

	for _, customer := range customersAPIResponse.Data {
		_, err := repo.GetCustomerByEmail(customer.Email)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newCustomer := &entities.Customer{
					FirstName: customer.FirstName,
					LastName:  customer.LastName,
					Email:     customer.Email,
					Avatar:    customer.Avatar,
				}
				_, err = repo.CreateCustomer(newCustomer)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// customer already exist, skip saving
		}
	}

	return nil
}
