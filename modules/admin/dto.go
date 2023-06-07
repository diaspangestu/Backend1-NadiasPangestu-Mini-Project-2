package admin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/entities"
)

type LoginAdminParam struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminParam struct {
	Username   string          `gorm:"column:username;unique"`
	RoleID     int             `gorm:"column:role_id"`
	IsVerified entities.Status `gorm:"column:is_verified;type enum('true','false')"`
	IsActived  entities.Status `gorm:"column:is_actived;type enum('true','false')"`
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

type SuccessCreate struct {
	dto.Response
	Data AdminParam `json:"data"`
}

type SuccessUpdate struct {
	dto.Response
	Data AdminParam `json:"data"`
}

type SuccessCreateCustomer struct {
	dto.Response
	Data CustomerParam `json:"data"`
}

type SuccessLoginAdmin struct {
	dto.Response
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SuccessGetAllCustomers struct {
	dto.Response
	Data interface{} `json:"data"`
}

type SuccessFetchCustomersFromAPI struct {
	dto.Response
}
