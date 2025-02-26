package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary Получить заказы
// @Description Возвращает список всех заказов
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/v1/orders [get]
func (h *Handlers) GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := h.s.GetAllOrders()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

// @Summary Получить заказ
// @Description Возвращает заказ по id
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /api/v1/orders/{id} [get]
func (h *Handlers) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		order, err := h.s.GetOrderById(id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}
