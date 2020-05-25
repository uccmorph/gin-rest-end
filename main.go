package main

import "github.com/gin-gonic/gin"

func handle1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	obj := NewResource()
	r := gin.Default()
	r.GET(obj.BuildHandler(API_GET_1))
	r.GET(obj.BuildHandler(API_GET_2))
	r.POST(obj.BuildHandler(API_POST_1))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
