// main.go
package main

import (
	"log"
	"temu-data/handlers"
	"temu-data/models"
	"temu-data/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "root:aabbcc1234@tcp(129.153.143.242:3306)/temu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Order{})
	return db
}

func setupRouter(orderHandler *handlers.OrderHandler) *gin.Engine {
	r := gin.Default()

	// API路由组
	api := r.Group("/api")
	{
		api.POST("/orders", orderHandler.CreateOrUpdate)
		api.GET("/orders/:sn", orderHandler.GetOrder)
	}

	return r
}

func main() {
	db := initDB()

	// 初始化服务和处理器
	orderService := services.NewOrderService(db)
	orderHandler := handlers.NewOrderHandler(orderService)

	// 设置路由
	r := setupRouter(orderHandler)

	// 启动服务器
	r.Run(":8080")
}
