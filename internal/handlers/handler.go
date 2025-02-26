package handlers

import (
	"github.com/gin-gonic/gin"
	_ "orders/cmd/docs"
	"orders/internal/services"
)

type Handlers struct {
	s services.OrderServiceInterface
}

func New(s services.OrderServiceInterface) *Handlers {
	return &Handlers{s: s}
}

func (h *Handlers) RegisterRoutes(r *gin.Engine) {
	r.GET("/api/v1/orders", h.GetOrders())
	r.GET("/api/v1/orders/:id", h.GetOrderById())
}
