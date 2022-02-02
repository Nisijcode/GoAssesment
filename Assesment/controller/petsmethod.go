package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pets struct {
	Id    int     `json:"id"`
	Type  string  `json:"type"`
	Price float32 `json:"price"`
}

func GetPetsHttp(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://petstore-demo-endpoint.execute-api.com/petstore/pets")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// return body
	petsArray := []Pets{}
	err = json.Unmarshal(body, &petsArray)
	if err != nil {
		fmt.Println(err.Error())
	}
	petsArrayByte, err := json.Marshal(petsArray)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(petsArrayByte)

}
func GetPets(c *gin.Context) {
	resp, err := http.Get("http://petstore-demo-endpoint.execute-api.com/petstore/pets")
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}

	// return body
	petsArray := []Pets{}
	err = json.Unmarshal(body, &petsArray)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
	}
	c.JSON(200, Response(Success, "", petsArray))

}
func PostPets(c *gin.Context) {
	var pet Pets
	err := c.BindJSON(&pet)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(BadRequestError, err.Error(), nil))
		return
	}
	resp, err := http.Get("http://petstore-demo-endpoint.execute-api.com/petstore/pets")
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}

	petsArray := []Pets{}

	err = json.Unmarshal(body, &petsArray)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}
	petsArray = append(petsArray, pet)
	c.JSON(200, Response(Success, "", petsArray))

}
func GetPetByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(BadRequestError, err.Error(), nil))
		return
	}
	resp, err := http.Get("http://petstore-demo-endpoint.execute-api.com/petstore/pets")
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}

	petsArray := []Pets{}

	err = json.Unmarshal(body, &petsArray)
	if err != nil {
		c.AbortWithStatusJSON(500, Response(Failed, err.Error(), nil))
		return
	}
	for _, a := range petsArray {
		if a.Id == idInt {
			c.IndentedJSON(Success, a)
			return
		} else {
			petsArray[0].Id = idInt
			c.IndentedJSON(Success, petsArray[0])
			return
		}
	}
	c.AbortWithStatusJSON(500, Response(NotFound, err.Error(), nil))
}
