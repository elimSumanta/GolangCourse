package delivery

import (
	"errors"

	serv "github.com/elim/GoCourses/Service"
)

type InitHTTPHandler struct {
	services *serv.Service
}

//New setup HTTP Handler
func New(serv *serv.Service) (handler *InitHTTPHandler, err error) {

	if serv == nil {
		err := errors.New("Service Not Found")
		return nil, err
	}

	http := &InitHTTPHandler{
		services: serv,
	}

	return http, nil
}
