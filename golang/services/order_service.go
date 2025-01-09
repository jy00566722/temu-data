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

func (s *OrderService) BatchCreateOrUpdateOrders(orders []*models.Order) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, order := range orders {
			var existingOrder models.Order
			result := tx.Where("sub_purchase_order_sn = ?", order.SubPurchaseOrderSn).First(&existingOrder)

			if err := tx.Where("sub_purchase_order_sn = ?", order.SubPurchaseOrderSn).Delete(&models.PackageDetail{}).Error; err != nil {
				return err
			}

			for _, detail := range order.PackageDetails {
				detail.SubPurchaseOrderSn = order.SubPurchaseOrderSn
				if err := tx.Create(&detail).Error; err != nil {
					return err
				}
			}

			if result.Error == gorm.ErrRecordNotFound {
				if err := tx.Create(order).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Model(&existingOrder).Updates(order).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *OrderService) GetOrderByPurchaseOrderSn(sn string) (*models.Order, error) {
	var order models.Order
	var details []models.PackageDetail

	if err := s.db.Where("sub_purchase_order_sn = ?", sn).First(&order).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("sub_purchase_order_sn = ?", sn).Find(&details).Error; err != nil {
		return nil, err
	}

	order.PackageDetails = details
	return &order, nil
}
