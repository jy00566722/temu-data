// services/order_service.go
package services

import (
	"temu-data/models"

	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrUpdateOrder(order *models.Order) error {
	var existingOrder models.Order
	result := s.db.Where("sub_purchase_order_sn = ?", order.SubPurchaseOrderSn).First(&existingOrder)

	if result.Error == gorm.ErrRecordNotFound {
		return s.db.Create(order).Error
	}

	return s.db.Model(&existingOrder).Updates(order).Error
}

func (s *OrderService) GetOrderByPurchaseOrderSn(sn string) (*models.Order, error) {
	var order models.Order
	result := s.db.Where("sub_purchase_order_sn = ?", sn).First(&order)
	return &order, result.Error
}
