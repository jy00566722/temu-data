// models/package_detail.go
package models

import (
	"gorm.io/gorm"
)

type PackageDetail struct {
	gorm.Model
	SubPurchaseOrderSn string `gorm:"index" json:"-"`
	ProductSkuId       int    `json:"productSkuId"`
	SkuNum             int    `json:"skuNum"`
}
