// models/package_detail.go
package models

import (
	"gorm.io/gorm"
)

type PackageDetail struct {
	gorm.Model
	OrderID          uint   `json:"-"`                // 外键关联到Order
	SkuId            string `json:"skuId"`            // SKU ID
	SkuCode          string `json:"skuCode"`          // SKU编码
	SkuName          string `json:"skuName"`          // SKU名称
	Color            string `json:"color"`            // 颜色
	Size             string `json:"size"`             // 尺码
	DeliveryQuantity int    `json:"deliveryQuantity"` // 发货数量
}
