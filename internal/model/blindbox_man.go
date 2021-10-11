package model

import (
	"github.com/dierbei/blind-box/pkg/forms"
	"gorm.io/gorm"
)

type Man struct {
	BaseModel
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

func (man *Man) Insert(tx *gorm.DB) error {
	if err := tx.Create(man).Error; err != nil {
		return err
	}
	return nil
}

func (man *Man) PageList(tx *gorm.DB, params *forms.ManListPageInput) ([]Man, int64, error) {
	list := make([]Man, 0)
	var count int64
	offset := (params.Page - 1) * params.PageSize

	result := tx.Limit(params.PageSize).Offset(offset).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, 0, result.Error
	}

	result = tx.Table(man.TableName()).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return list, count, nil
}

func (man *Man) SelectByUserID(tx *gorm.DB, id int64) ([]Man, error) {
	list := make([]Man, 0)

	result := tx.Where("user_id = ?", id).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	return list, nil
}

func (Man) TableName() string {
	return "blind_box_man"
}
