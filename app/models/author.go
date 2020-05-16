package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name     string
	NameKana string
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
	db.Table("authors").Order("name_kana asc").Where("name LIKE ? OR name_kana LIKE ? ", name, name).Select("id, name, name_kana").Scan(&authors)
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
	db.Table("records").Group("authors.name").Order("count desc").Select("authors.name, count(*) as count").Joins("left join authors on records.author = authors.id").Limit(10).Scan(&authorsCount)
	return authorsCount
}
