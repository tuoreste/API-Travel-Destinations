package routes

import (
    "github.com/gin-gonic/gin"
)

func SetupDestinationRoutes(router *gin.Engine) {
    router.GET("/destinations", controllers.GetDestinations)
    router.GET("/destinations/:id", controllers.GetDestinationByID)
    router.POST("/destinations", controllers.CreateDestination)
}
