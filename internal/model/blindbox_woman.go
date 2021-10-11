package model

import (
	"github.com/dierbei/blind-box/pkg/forms"
	"gorm.io/gorm"
)

type Woman struct {
	BaseModel
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

func (woman *Woman) Insert(tx *gorm.DB) error {
	if err := tx.Create(woman).Error; err != nil {
		return err
	}
	return nil
}

func (woman *Woman) PageList(tx *gorm.DB, params *forms.WomanListPageInput) ([]Woman, int64, error) {
	list := make([]Woman, 0)
	var count int64
	offset := (params.Page - 1) * params.PageSize

	result := tx.Limit(params.PageSize).Offset(offset).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, 0, result.Error
	}

	result = tx.Table(woman.TableName()).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return list, count, nil
}

func (woman *Woman) SelectByUserID(tx *gorm.DB, id int64) ([]Man, error) {
	list := make([]Man, 0)

	result := tx.Where("user_id = ?", id).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	return list, nil
}

func (Woman) TableName() string {
	return "blind_box_woman"
}
