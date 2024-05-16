package userHandlers

import userControllers "github.com/guivictorr/fast-feet-go/controllers/user-controllers"

type handler struct {
	service userControllers.Service
}

func NewHandler(service userControllers.Service) *handler {
	return &handler{service: service}
}
