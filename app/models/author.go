package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name     string `json:"name"`
	NameKana string `json:"name_kana"`
}

type AuthorCount struct {
	gorm.Model
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func ApiFindAuthor(id string) Author {
	var author Author
	db.First(&author, id)
	return author
}

func ApiFindAuthorByName(param string) []Author {
	var authors []Author
	name := fmt.Sprintf("%%%s%%", param)
	db.Table("authors").
		Select("id, name, name_kana").
		Where("name LIKE ? OR name_kana LIKE ? ", name, name).
		Order("name_kana asc").
		Scan(&authors)
	return authors
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

func ApiFindCountByAuthor() []AuthorCount {
	var authorsCount []AuthorCount
	db.Table("records").
		Joins("left join authors on records.author = authors.id").
		Select("authors.name, count(*) as count").
		Group("authors.name").
		Order("count desc").
		Limit(10).Scan(&authorsCount)
	return authorsCount
}
