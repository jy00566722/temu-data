// models/order.go
package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	SubPurchaseOrderSn string `gorm:"uniqueIndex;type:varchar(100)" json:"subPurchaseOrderSn"`
	DeliveryOrderSn    string `json:"deliveryOrderSn"`
	ExpressCompany     string `json:"expressCompany"`
	ExpressDeliverySn  string `json:"expressDeliverySn"`
	SubWarehouseName   string `json:"subWarehouseName"`
	SupplierId         string `json:"supplierId"`
	ProductSkcId       string `json:"productSkcId"`
	ProductName        string `json:"productName"`
	ProductSkcPicture  string `json:"productSkcPicture"`
	SkcExtCode         string `json:"skcExtCode"`
	DeliverPackageNum  int    `json:"deliverPackageNum"`
	DeliverSkcNum      int    `json:"deliverSkcNum"`
	// PackageDetailList  string `gorm:"type:text" json:"packageDetailList"`
	PackageDetails []PackageDetail `gorm:"foreignKey:OrderID" json:"packageDetailList"`
	Status         string          `json:"status"`
}
