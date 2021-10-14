package model

import (
	"gorm.io/gorm"
)

// People And Image：一对多
// 添加自拍的时候需要添加PeopleID
type Image struct {
	ID       int32  `json:"id"`
	PeopleID int32  `json:"people_id"`
	Url      string `json:"url" gorm:"type:blob"`
}

func (model *Image) Insert(tx *gorm.DB, images []Image) error {
	if err := tx.Create(&images).Error; err != nil {
		return err
	}
	return nil
}

func (model *Image) SelectByPeopleID(tx *gorm.DB) ([]Image, error) {
	images := make([]Image, 0)
	if err := tx.Where("people_id = ?", &model.PeopleID).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (Image) TableName() string {
	return "blind_box_image"
}
