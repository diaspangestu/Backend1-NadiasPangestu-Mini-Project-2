package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
)

type ControllerSuperadminInterface interface {
	CreateCustomer(req CustomerParam) (interface{}, error)
	ApprovedAdminRegister(id uint) (interface{}, error)
	RejectedAdminRegister(id uint) (interface{}, error)
}

type ControllerSuperadmin struct {
	uc UsecaseSuperadminInterface
}

// CreateCustomer Superadmin
func (ctrl ControllerSuperadmin) CreateCustomer(req CustomerParam) (interface{}, error) {
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

func (ctrl ControllerSuperadmin) ApprovedAdminRegister(id uint) (interface{}, error) {
	err := ctrl.uc.ApprovedAdminRegister(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Approved Registration Admin",
			Message:      "Approved",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) RejectedAdminRegister(id uint) (interface{}, error) {
	err := ctrl.uc.RejectedAdminRegister(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Rejected Registration Admin",
			Message:      "Rejected",
			ResponseTime: "",
		},
	}

	return response, nil
}
