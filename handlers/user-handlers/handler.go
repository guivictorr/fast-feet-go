package userHandlers

import authControllers "github.com/guivictorr/fast-feet-go/controllers/auth-controllers"

type handler struct {
	service authControllers.Service
}

func NewHandler(service authControllers.Service) *handler {
	return &handler{service: service}
}
