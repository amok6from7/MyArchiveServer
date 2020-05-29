package models

import (
	"github.com/jinzhu/gorm"
)

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
	err := db.Limit(5).Order("created_at desc").Find(&asyncManage).Error
	OutputLogIfError(err)
	return asyncManage
}

func CreateAsyncManage(actionName string) uint {
	asyncManage := AsyncManage{
		Action: actionName,
		Status: START,
	}
	err := db.Create(&asyncManage).Error
	OutputLogIfError(err)
	return asyncManage.ID
}

func UpdateAsyncManage(id uint) {
	var asyncManage AsyncManage
	err := db.First(&asyncManage, id).Error
	OutputLogIfError(err)
	asyncManage.Status = COMPLETE
	err = db.Save(&asyncManage).Error
	OutputLogIfError(err)
}
