package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var message string

func root(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"message": message,
	})
}

func main() {
	log.Println("Starting")
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", root)
	r.Run(":80")
}
