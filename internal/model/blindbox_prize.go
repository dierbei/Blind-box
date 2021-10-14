package model

import "gorm.io/gorm"

type UserPrize struct {
	UserID  int32 `json:"user_id"`
	PrizeID int32 `json:"prize_id"`
}

func (model *UserPrize) Exist(tx *gorm.DB) bool {
	userPrize := UserPrize{}
	result := tx.Where("user_id = ? and prize_id = ?", model.UserID, model.PrizeID).First(&userPrize)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (userPrize *UserPrize) SelectList(tx *gorm.DB) ([]UserPrize, error) {
	userPrizes := make([]UserPrize, 0)
	result := tx.Where("user_id = ?", userPrize.UserID).Find(&userPrizes)
	if result.Error != nil {
		return nil, result.Error
	}
	return userPrizes, nil
}

func (userPrize *UserPrize) Insert(tx *gorm.DB) error {
	if err := tx.Create(&userPrize).Error; err != nil {
		return err
	}
	return nil
}

func (UserPrize) TableName() string {
	return "blind_box_user_prize"
}
