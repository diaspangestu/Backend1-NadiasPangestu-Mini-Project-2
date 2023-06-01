package user

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type Controller struct {
	uc UsecaseInterface
}

type ControllerInterface interface {
	GetUserByID(payload Payload) dto.Response
}
