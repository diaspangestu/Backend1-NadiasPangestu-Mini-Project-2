package admin

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type ControllerAdminInterface interface {
	LoginAdmin(username, password string) (interface{}, error)
	RegisterAdmin(req AdminParam) (interface{}, error)
	CreateCustomer(req CustomerParam) (interface{}, error)
	DeleteCustomerById(id uint) error
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) (interface{}, error)
}

type ControllerAdmin struct {
	uc UsecaseAdmin
}

func (ctrl ControllerAdmin) LoginAdmin(username, password string) (interface{}, error) {
	admin, err := ctrl.uc.LoginAdmin(username, password)
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
	}

	return response, nil
}

func (ctrl ControllerAdmin) RegisterAdmin(req AdminParam) (interface{}, error) {
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
