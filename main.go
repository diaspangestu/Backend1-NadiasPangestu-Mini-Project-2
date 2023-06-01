package main

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/user"
)

func main() {
	var request = dto.Request{
		Body: map[string]string{
			"id": "1",
		},
		Method: "GET",
		Path:   "/get-user",
		Header: map[string]string{
			"Authorization": "token",
		},
	}

	router := user.NewRouter()
	router.Route(request)

}
