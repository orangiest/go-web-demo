package main

import (
	"cwm.wiki/web/controller"
	"cwm.wiki/web/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books",controller.FindBooks)
	r.POST("/books",controller.CreateBook)
	r.GET("/books/:id",controller.FindBook)
	r.PATCH("/books/:id",controller.UpdateBook)
	r.DELETE("/books/:id",controller.DeleteBooks)

	r.Run()



}
