package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test/manager"
	"test/models"
)

func FindPet(c *gin.Context) {
	var pet models.Pet
	id := c.Param("id")
	res := manager.GetPetById(id)
	copier.Copy(&pet, &res)
	c.JSON(http.StatusOK, gin.H{"data": pet})
}
