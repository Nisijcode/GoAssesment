package main

import (
	// controller "NisijAssesment/controller"
	handlers "NisijAssesment/handlers"

	"fmt"
	"log"
	"net/http"
)

// import (
// 	controller "NisijAssesment/controller"

// 	"github.com/gin-gonic/gin"
// )

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":8989"),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}

	// router := gin.Default()
	// router.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}
	// 	c.Next()
	// })
	// router.GET("/pets", controller.GetPets)
	// router.POST("/pets", controller.PostPets)
	// router.GET("/pets/:id", controller.GetPetByID)
	// router.Run(":" + "8989")
}
