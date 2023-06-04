package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
)

type ControllerSuperadminInterface interface {
	LoginSuperadmin(username, password string) (interface{}, error)
	CreateCustomer(req CustomerParam) (interface{}, error)
	DeleteCustomerById(id uint) error
	ApprovedAdminRegister(id uint) (interface{}, error)
	RejectedAdminRegister(id uint) (interface{}, error)
	UpdateActivedAdmin(id uint) (interface{}, error)
	UpdateDeadactivedAdmin(id uint) (interface{}, error)
	GetApprovalRequest() (interface{}, error)
}

type ControllerSuperadmin struct {
	uc UsecaseSuperadminInterface
}

func (ctrl ControllerSuperadmin) LoginSuperadmin(username, password string) (interface{}, error) {
	superAdmin, err := ctrl.uc.LoginSuperadmin(username, password)
	if err != nil {
		return nil, err
	}

	response := SuccessLoginSuperadmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Login Successfull",
			Message:      "Success",
			ResponseTime: "",
		},
		Username: superAdmin.Username,
	}

	return response, nil
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

// DeleteCustomerById Superadmin
func (ctrl ControllerSuperadmin) DeleteCustomerById(id uint) error {
	err := ctrl.uc.DeleteCustomerById(id)
	if err != nil {
		return err
	}

	return nil
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

func (ctrl ControllerSuperadmin) UpdateActivedAdmin(id uint) (interface{}, error) {
	err := ctrl.uc.UpdateActivedAdmin(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Admin is Actived Now",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) UpdateDeadactivedAdmin(id uint) (interface{}, error) {
	err := ctrl.uc.UpdateDeadactivedAdmin(id)
	if err != nil {
		return SuccessUpdateRegisterAdmin{}, err
	}

	response := SuccessUpdateRegisterAdmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Admin is Deadactived Now",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) GetApprovalRequest() (interface{}, error) {
	request, err := ctrl.uc.GetApprovalRequest()
	if err != nil {
		return SuccessGetApprovalRequest{}, err
	}

	response := SuccessGetApprovalRequest{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Get All Data Approval Request",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: request,
	}

	return response, nil
}
