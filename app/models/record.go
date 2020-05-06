package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type RECORD struct {
	gorm.Model
	Title      string
	TitleKana  string
	Evaluation string
	Author     int
}

type Result struct {
	gorm.Model
	Title string `json:"title"`
	Name  string `json:"name"`
	Evaluation string `json:"evaluation"`
}

func FindById(id string) []Result {
	var results []Result
	db.Table("records").Where("records.ID = ?", id).Select("records.title, authors.name, records.evaluation").Joins("left join authors on authors.id = records.author").Scan(&results)
	return results//TODO order by author name asc
}

func FindByTitle(param string) []Result {
	var results []Result
	title := fmt.Sprintf("%%%s%%", param)
	db.Table("records").Where("records.title LIKE ? OR records.title_kana LIKE ?", title, title).Select("records.title, authors.name, records.evaluation").Joins("left join authors on authors.id = records.author").Scan(&results)
	return results//TODO order by author name asc
}

func FindByAuthor(param string) []Result {
	var results []Result
	name := fmt.Sprintf("%%%s%%", param)
	db.Table("records").Where("authors.name LIKE ? OR authors.name_kana LIKE ? ", name, name).Select("records.title, authors.name, records.evaluation").Joins("left join authors on authors.id = records.author").Scan(&results)
	return results
}

func CreateRecord(record *RECORD) {
	db.Create(&record)
}

func UpdateRecord(id string, params *RECORD) {
	var record RECORD
	db.First(&record, id)
	record.Title = params.Title
	record.TitleKana = params.TitleKana
	record.Evaluation = params.Evaluation
	record.Author = params.Author
	db.Save(&record)
}

func DeleteRecord(id string) {
	var record RECORD
	db.First(&record, id)
	db.Delete(&record)
}