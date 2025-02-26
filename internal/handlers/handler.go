package handlers

import "orders/internal/services"

type Handlers struct {
	s services.OrderServiceInterface
}
