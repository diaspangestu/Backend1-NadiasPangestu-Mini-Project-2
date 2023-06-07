package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
)

type ControllerSuperadminInterface interface {
	CreateSuperadmin(req SuperAdminParam) (interface{}, error)
	LoginSuperadmin(id uint, username, password string) (interface{}, error)
	CreateCustomer(req CustomerParam) (interface{}, error)
	DeleteCustomerById(id uint) error
	GetAllCustomers(first_name, last_name, email string, page, pageSize int) (interface{}, error)
	ApprovedAdminRegister(id uint) (interface{}, error)
	RejectedAdminRegister(id uint) (interface{}, error)
	UpdateActivedAdmin(id uint) (interface{}, error)
	UpdateDeadactivedAdmin(id uint) (interface{}, error)
	GetApprovalRequest() (interface{}, error)
	GetAllAdmins(username string, page, pageSize int) (interface{}, error)
}

type ControllerSuperadmin struct {
	uc UsecaseSuperadminInterface
}

func (ctrl ControllerSuperadmin) CreateSuperadmin(req SuperAdminParam) (interface{}, error) {
	_, err := ctrl.uc.CreateSuperadmin(req)
	if err != nil {
		return SuccessCreateSuperadmin{}, err
	}

	response := SuccessCreateSuperadmin{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Create Superadmin",
			Message:      "Success",
			ResponseTime: "",
		},
	}

	return response, nil
}

func (ctrl ControllerSuperadmin) LoginSuperadmin(id uint, username, password string) (interface{}, error) {
	superAdmin, tokenString, err := ctrl.uc.LoginSuperadmin(id, username, password)
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
		Token:    tokenString,
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
			MessageTitle: "Success Create Data",
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

// GetAllCustomers Superadmin
func (ctrl ControllerSuperadmin) GetAllCustomers(first_name, last_name, email string, page, pageSize int) (interface{}, error) {
	request, err := ctrl.uc.GetAllCustomers(first_name, last_name, email, page, pageSize)
	if err != nil {
		return SuccessGetAllCustomers{}, err
	}

	response := SuccessGetAllCustomers{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get All Data Data",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: request,
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

func (ctrl ControllerSuperadmin) GetAllAdmins(username string, page, pageSize int) (interface{}, error) {
	request, err := ctrl.uc.GetAllAdmins(username, page, pageSize)
	if err != nil {
		return SuccessGetAllAdmins{}, err
	}

	response := SuccessGetAllAdmins{
		Response: dto.Response{
			Success:      true,
			MessageTitle: "Success Get All Admin Data",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: request,
	}

	return response, nil
}
