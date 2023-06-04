package superadmin

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type CustomerParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type LoginSuperadmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SuccessCreate struct {
	dto.Response
	Data CustomerParam `json:"data"`
}

type SuccessUpdateRegisterAdmin struct {
	dto.Response
}

type SuccessGetApprovalRequest struct {
	dto.Response
	Data interface{} `json:"data"`
}

type SuccessLoginSuperadmin struct {
	dto.Response
	Username string `json:"username"`
}

type SuccessGetAllCustomers struct {
	dto.Response
	Data interface{} `json:"data"`
}
