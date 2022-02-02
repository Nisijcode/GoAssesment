package controllers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func SavePet(c *gin.Context) {

	var req models.SavePet
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pets := models.SavePet{Type: req.Type, Price: req.Price}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}
