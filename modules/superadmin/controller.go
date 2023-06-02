package superadmin

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type ControllerRegisterApprovalInterface interface {
	ApprovedAdminRegister(id uint) (interface{}, error)
	RejectedAdminRegister(id uint) (interface{}, error)
}

type ControllerRegisterApproval struct {
	uc UsecaseSuperadminInterface
}

func (ctrl ControllerRegisterApproval) ApprovedAdminRegister(id uint) (interface{}, error) {
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

func (ctrl ControllerRegisterApproval) RejectedAdminRegister(id uint) (interface{}, error) {
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
