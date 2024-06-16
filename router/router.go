package router

import (
    "github.com/gin-gonic/gin"
    "github.com/philipphahmann/rating-app/handlers"
)

// SetupRouter sets up the Gin router
func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Define the route
    r.GET("/record/:id", handlers.GetRecordByID)
    r.GET("/records", handlers.GetAllRecords)
    r.PATCH("/record/:id", handlers.UpdateRecordByID)
    r.DELETE("/record/:id", handlers.DeleteRecordByID)
    r.POST("/record", handlers.CreateRecordByID)

    return r
}
