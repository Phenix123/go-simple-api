package services

import (
	"orders/internal/models"
	"orders/internal/repositories"
)

type OrderServiceInterface interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderById(id int64) (*models.Order, error)
}

type OrderService struct {
	r repositories.OrderRepositoryInterface
}

func New(r repositories.OrderRepositoryInterface) *OrderService {
	return &OrderService{r: r}
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.r.GetOrders()
}
func (s *OrderService) GetOrderById(id int64) (*models.Order, error) {
	return s.r.GetOrderById(id)
}
