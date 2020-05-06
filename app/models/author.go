package models

import (
	"github.com/jinzhu/gorm"
)

type AUTHOR struct {
	gorm.Model
	Name     string
	NameKana string
}

func ApiFindAuthor(id string) AUTHOR {
	var author AUTHOR
	db.First(&author, id)
	return author
}

func CreateAuthor(author *AUTHOR) {
	db.Create(&author)
}

func UpdateAuthor(id string, params *AUTHOR) {
	var author AUTHOR
	db.First(&author, id)
	author.Name = params.Name
	author.NameKana = params.NameKana
	db.Save(&author)
}

func DeleteAuthor(id string) {
	var author AUTHOR
	db.First(&author, id)
	db.Delete(&author)
}