package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)
func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:orders_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:orders_id", controller.UpdateOrder())
}