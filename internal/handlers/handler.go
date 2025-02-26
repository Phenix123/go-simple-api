package handlers

import (
	"orders/internal/services"
)

type Handlers struct {
	s services.OrderServiceInterface
}

func New(s services.OrderServiceInterface) *Handlers {
	return &Handlers{s: s}
}
