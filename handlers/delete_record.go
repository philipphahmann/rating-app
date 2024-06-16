package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// HelloHandler handles the /hello route
func DeleteRecordByID(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Record deleted",
    })
}
