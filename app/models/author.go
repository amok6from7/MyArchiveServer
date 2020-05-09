package models

import (
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name     string
	NameKana string
}

func ApiFindAuthor(id string) Author {
	var author Author
	db.First(&author, id)
	return author
}

func CreateAuthor(author *Author) {
	db.Create(&author)
}

func UpdateAuthor(id string, params *Author) {
	var author Author
	db.First(&author, id)
	author.Name = params.Name
	author.NameKana = params.NameKana
	db.Save(&author)
}

func DeleteAuthor(id string) {
	var author Author
	db.First(&author, id)
	db.Delete(&author)
}
