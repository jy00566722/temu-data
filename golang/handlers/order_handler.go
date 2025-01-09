// handlers/order_handler.go
package handlers

import (
	"temu-data/models"
	"temu-data/services"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) BatchCreateOrUpdate(c *gin.Context) {
	var orders []*models.Order
	if err := c.ShouldBindJSON(&orders); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.BatchCreateOrUpdateOrders(orders); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"count":   len(orders),
	})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	sn := c.Param("sn")
	order, err := h.orderService.GetOrderByPurchaseOrderSn(sn)
	if err != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	c.JSON(200, order)
}
