package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test/manager"
	"test/models"
)

func FindAllPets(c *gin.Context) {
	pet := make([]models.Pet, 0)
	petInfo := manager.GetAllPets()
	copier.Copy(&pet, &petInfo)
	c.JSON(http.StatusOK, gin.H{"data": pet})
}
