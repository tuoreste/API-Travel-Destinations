package routes

import (
    "travel-destinations-api/controllers"
    "github.com/gin-gonic/gin"
)

func SetupReviewRoutes(router *gin.Engine) {
    router.POST("/reviews", controllers.SubmitReview)
    router.GET("/reviews/:destination_id", controllers.GetReviewsByDestination)
}
