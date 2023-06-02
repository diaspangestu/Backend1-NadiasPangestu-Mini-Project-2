package admin

import "github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/dto"

type AdminParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SuccessCreate struct {
	dto.Response
	Data AdminParam `json:"data"`
}
