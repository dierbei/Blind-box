package model

import "gorm.io/gorm"

type User struct {
	BaseModel
	Openid string `json:"openid"`
}

func (model *User) SelectAll(tx *gorm.DB) ([]User, error) {
	var user []User
	if err := tx.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (model *User) Select(tx *gorm.DB) (*User, error) {
	user := User{}
	if err := tx.Where("openid = ?", model.Openid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (user *User) Insert(tx *gorm.DB) error {
	if err := tx.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) ExistOpenID(tx *gorm.DB) bool {
	result := tx.Where("openid = ?", user.Openid).First(&User{})
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (User) TableName() string {
	return "blind_box_user"
}
