package admin

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type ControllerAdminInterface interface {
	// admin
	LoginAdmin(username, password string) (interface{}, error)
	RegisterAdmin(req LoginAdminParam) (interface{}, error)
	GetAdminById(id uint) (interface{}, error)
	UpdateAdminById(id uint, admin AdminParam) (interface{}, error)
	DeleteAdminById(id uint) error
	// customer
	CreateCustomer(req CustomerParam) (interface{}, error)
	DeleteCustomerById(id uint) error
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) (interface{}, error)
	SaveCustomersFromAPI() (interface{}, error)
}

type ControllerAdmin struct {
	uc UsecaseAdmin
}

func (ctrl ControllerAdmin) LoginAdmin(id uint, username, password string) (interface{}, error) {
	admin, tokenString, err := ctrl.uc.LoginAdmin(id, username, password)
	if err != nil {
		return nil, err
	}

	response := SuccessLoginAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Login Successful",
			Message:      "Sucess",
			ResponseTime: "",
		},
		Username: admin.Username,
		Token:    tokenString,
	}

	return response, nil
}

func (ctrl ControllerAdmin) RegisterAdmin(req LoginAdminParam) (interface{}, error) {
	_, err := ctrl.uc.RegisterAdmin(req)
	if err != nil {
		return SuccessCreateAdmin{}, err
	}

	response := SuccessCreateAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Register Admin",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerAdmin) GetAdminById(id uint) (interface{}, error) {
	admin, err := ctrl.uc.GetAdminById(id)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessCreate{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get Admin",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username: admin.Username,
			Password: admin.Password,
		},
	}

	return response, nil
}

func (ctrl ControllerAdmin) UpdateAdminById(id uint, admin AdminParam) (interface{}, error) {
	updatedAdmin, err := ctrl.uc.UpdateCustomerById(id, admin)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessUpdate{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Update Admin",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: AdminParam{
			Username: updatedAdmin.Username,
			Password: updatedAdmin.Password,
		},
	}

	return response, nil
}

func (ctrl ControllerAdmin) DeleteAdminById(id uint) error {
	err := ctrl.uc.DeleteAdminById(id)
	if err != nil {
		return err
	}

	return nil
}

// CreateCustomer Admin
func (ctrl ControllerAdmin) CreateCustomer(req CustomerParam) (interface{}, error) {
	customer, err := ctrl.uc.CreateCustomer(req)
	if err != nil {
		return SuccessCreateCustomer{}, err
	}

	response := SuccessCreateCustomer{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Create Customer",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}

	return response, nil
}

// DeleteCustomerById Admin
func (ctrl ControllerAdmin) DeleteCustomerById(id uint) error {
	err := ctrl.uc.DeleteCustomerById(id)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl ControllerAdmin) GetAllCustomers(first_name, last_name, email string, page, pageSize int) (interface{}, error) {
	request, err := ctrl.uc.GetAllCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		return SuccessGetAllCustomers{}, err
	}

	response := SuccessGetAllCustomers{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get All Customer Data",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: request,
	}

	return response, nil
}

func (ctrl ControllerAdmin) SaveCustomersFromAPI() (interface{}, error) {
	err := ctrl.uc.SaveCustomersFromAPI()
	if err != nil {
		return SuccessFetchCustomersFromAPI{}, err
	}

	response := SuccessFetchCustomersFromAPI{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Fetch Customer Data from API",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}
