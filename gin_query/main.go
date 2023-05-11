package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(ctx *gin.Context) {
		a := ctx.DefaultQuery("a", "2")
		b := ctx.Query("b")

		ctx.String(http.StatusOK, "query is %s %s", a, b)
	})
	router.POST("/post", func(ctx *gin.Context) {
		id := ctx.Query("id")

		name := ctx.PostForm("name")
		fmt.Println(name)
		ctx.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
		})
	})
	router.Run()
}
