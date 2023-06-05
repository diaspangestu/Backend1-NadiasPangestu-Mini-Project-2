package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
)

type AdminParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomerParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type SuccessCreateAdmin struct {
	dto.Response
}

type SuccessCreateCustomer struct {
	dto.Response
	Data CustomerParam `json:"data"`
}

type SuccessLoginAdmin struct {
	dto.Response
	Username string `json:"username"`
}

type SuccessGetAllCustomers struct {
	dto.Response
	Data interface{} `json:"data"`
}
