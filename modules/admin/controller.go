package admin

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type ControllerAdminInterface interface {
	LoginAdmin(username, password string) (interface{}, error)
	CreateCustomer(req CustomerParam) (interface{}, error)
	RegisterAdmin(req AdminParam) (interface{}, error)
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
			MessageTitle: "Login Successfull",
			Message:      "Sucess",
			ResponseTime: "",
		},
		Username: admin.Username,
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

func (ctrl ControllerAdmin) RegisterAdmin(req AdminParam) (interface{}, error) {
	admin, err := ctrl.uc.RegisterAdmin(req)
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
		Data: AdminParam{
			Username: admin.Username,
			Password: admin.Password,
		},
	}

	return response, nil
}
