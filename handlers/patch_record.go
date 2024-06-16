package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// HelloHandler handles the /hello route
func UpdateRecordByID(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Record updated successfully",
    })
}
