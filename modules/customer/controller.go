package customer

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type ControllerCustomerInterface interface {
	CreateCustomer(req CustomerParam) (interface{}, error)
	GetCustomerById(id uint) (interface{}, error)
	UpdateCustomerById(id uint, customer CustomerParam) (interface{}, error)
}

type ControllerCustomer struct {
	uc UsecaseCustomer
}

func (ctrl ControllerCustomer) CreateCustomer(req CustomerParam) (interface{}, error) {
	customer, err := ctrl.uc.CreateCustomer(req)
	if err != nil {
		return SuccessCreate{}, err
	}

	response := SuccessCreate{
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

func (ctrl ControllerCustomer) GetCustomerById(id uint) (interface{}, error) {
	customer, err := ctrl.uc.GetCustomerById(id)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessCreate{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get Customer",
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

func (ctrl ControllerCustomer) UpdateCustomerById(id uint, customer CustomerParam) (interface{}, error) {
	updatedCustomer, err := ctrl.uc.UpdateCustomerById(id, customer)
	if err != nil {
		return dto.ErrorResponse{}, err
	}

	response := SuccessUpdate{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Update Customer",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: updatedCustomer.FirstName,
			LastName:  updatedCustomer.LastName,
			Email:     updatedCustomer.Email,
			Avatar:    updatedCustomer.Email,
		},
	}

	return response, nil
}
