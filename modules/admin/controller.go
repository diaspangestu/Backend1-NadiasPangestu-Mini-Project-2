package admin

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type ControllerAdminInterface interface {
	RegisterAdmin(req AdminParam) (interface{}, error)
}

type ControllerAdmin struct {
	uc UsecaseAdmin
}

func (ctrl ControllerAdmin) RegisterAdmin(req AdminParam) (interface{}, error) {
	admin, err := ctrl.uc.RegisterAdmin(req)
	if err != nil {
		return SuccessCreate{}, err
	}

	response := SuccessCreate{
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
