package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler Handler) PostCSV(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "./uploaded.csv")
	c.JSON(http.StatusOK, gin.H{"message": "file uploaded"})
}
