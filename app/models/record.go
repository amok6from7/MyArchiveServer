package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Record struct {
	gorm.Model
	Title      string
	TitleKana  string
	Evaluation string
	Author     int
}

type Result struct {
	gorm.Model
	Title      string `json:"title"`
	TitleKana  string `json:"title_kana"`
	Name       string `json:"name"`
	Evaluation string `json:"evaluation"`
	AuthorId   int    `json:"author_id"`
}

func FindById(id string) []Result {
	var results []Result
	db.Table("records").
		Joins("left join authors on authors.id = records.author").
		Select("records.id, records.title, records.title_kana, records.evaluation, authors.id as author_id, authors.name").
		Where("records.ID = ? AND records.deleted_at IS NULL", id).
		Scan(&results)
	return results
}

func FindByTitle(param string) []Result {
	var results []Result
	title := fmt.Sprintf("%%%s%%", param)
	db.Table("records").
		Joins("left join authors on authors.id = records.author").
		Select("records.id, records.title, records.evaluation, authors.id as author_id, authors.name").
		Where("(records.title LIKE ? OR records.title_kana LIKE ?) AND records.deleted_at IS NULL", title, title).
		Order("authors.name_kana COLLATE \"ja_JP.utf8\" asc").
		Scan(&results)
	return results
}

func FindByAuthor(param string) []Result {
	var results []Result
	name := fmt.Sprintf("%%%s%%", param)
	db.Table("records").
		Joins("left join authors on authors.id = records.author").
		Select("records.id, records.title, authors.id as author_id, authors.name, records.evaluation").
		Where("(authors.name LIKE ? OR authors.name_kana LIKE ?) AND records.deleted_at IS NULL", name, name).
		Order("authors.name_kana COLLATE \"ja_JP.utf8\" asc, records.title_kana COLLATE \"ja_JP.utf8\" asc").
		Scan(&results)
	return results
}

func CreateRecord(record *Record) {
	db.Create(&record)
}

func UpdateRecord(id string, params *Record) {
	var record Record
	db.First(&record, id)
	record.Title = params.Title
	record.TitleKana = params.TitleKana
	record.Evaluation = params.Evaluation
	record.Author = params.Author
	db.Save(&record)
}

func DeleteRecord(id string) {
	var record Record
	db.First(&record, id)
	db.Delete(&record)
}
