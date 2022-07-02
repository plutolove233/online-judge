package main

import "github.com/gin-gonic/gin"

func GoTest(c *gin.Context)  {
	c.JSON(200, gin.H{
		"code": 2000,
		"message": "Hello Go!",
	})
}

func main() {
	engine  := gin.Default()
	router := engine.Group("test")
	router.GET("hello", GoTest)
	err := engine.Run("0.0.0.0:8080")
	if err != nil {
		println(err)
	}
}
