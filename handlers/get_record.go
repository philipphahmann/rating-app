package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// HelloHandler handles the /hello route
func GetRecordByID(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Record found successfully",
    })
}
