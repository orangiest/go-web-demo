package mapper

import "cwm.wiki/web/models"

func SelectBooks() []models.Book{
	var books []models.Book

	models.DB.Find(&books)

	return books
}