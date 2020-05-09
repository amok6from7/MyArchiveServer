package models

import "github.com/jinzhu/gorm"

type Status int

const (
	START Status = iota
	COMPLETE
)

type AsyncManage struct {
	gorm.Model
	Action string
	Status Status
}

func FindAsyncManage() []AsyncManage {
	var asyncManage []AsyncManage
	db.Limit(5).Order("created_at desc").Find(&asyncManage)
	return asyncManage
}

func CreateAsyncManage(actionName string) uint {
	asyncManage := AsyncManage{
		Action: actionName,
		Status: START,
	}
	db.Create(&asyncManage)
	return asyncManage.ID
}

func UpdateAsyncManage(id uint) {
	var asyncManage AsyncManage
	db.First(&asyncManage, id)
	asyncManage.Status = COMPLETE
	db.Save(&asyncManage)
}
