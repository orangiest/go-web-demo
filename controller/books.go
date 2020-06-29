package controller

import (
	"cwm.wiki/web/mapper"
	"cwm.wiki/web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// get
func FindBooks(c *gin.Context) {
	books := mapper.SelectBooks()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// post
func CreateBook(c *gin.Context) {

	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// findABook
func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// updateBook
func UpdateBook(c *gin.Context) {
	// 这些controller都没有返回值
	var book models.Book
	if err := models.DB.Where("id = ?",c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Record not found"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Update(input)

	c.JSON(http.StatusOK,gin.H{"data": book})

}


// delete

func DeleteBooks(c *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"err":err.Error() + " Record not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK,gin.H{"data": true})

}
