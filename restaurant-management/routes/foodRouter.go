package routes

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)

func foodRoutes(incomingRoutes *gin.engine){
	incomingRoutes.GET("/foods", controller.Getfoods()))
	incomingRoutes.GET
	incomingRoutes.POST
	incomingRoutes.PATCH
}