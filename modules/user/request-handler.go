package user

import (
	"fmt"

	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

type RequestHandlerInterface interface {
	GetUserByID(request dto.Request) dto.Response
}

func (rq RequestHandler) GetUserByID(request dto.Request) dto.Response {

	// convert response ke payload, terjadi validasi
	payload := Payload{
		ID: 1,
	}

	response := rq.ctrl.GetUserByID(payload)

	fmt.Println(response)
	return response
}
