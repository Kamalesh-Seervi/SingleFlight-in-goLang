package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ServerSetup() {
	r := gin.Default()

	r.GET("/api/v1/get_something", func(c *gin.Context) {
		name := c.DefaultQuery("name", "")
		response := processingRequest(name)
		c.String(200, response)
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func processingRequest(name string) string {
	fmt.Println("[DEBUG] processing request..")
	return "Hi there! You requested " + name
}
