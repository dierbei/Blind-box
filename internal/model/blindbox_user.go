package model

import "gorm.io/gorm"

type User struct {
	BaseModel
	Openid string `json:"openid"`
	Status int    `json:"status" gorm:"default:0"`
}

func (userModel *User) UpdateStatus(tx *gorm.DB) error {
	user := User{}
	if err := tx.Table(userModel.TableName()).Where("id = ?", userModel.ID).First(&user).Error; err != nil {
		return err
	}
	if user.Status != 0 {
		return nil
	}

	if err := tx.Table(userModel.TableName()).Where("id = ?", userModel.ID).Update("status", 1).Error; err != nil {
		return err
	}
	return nil
}

func (userModel *User) SelectAll(tx *gorm.DB) ([]User, error) {
	var user []User
	if err := tx.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (userModel *User) SelectByID(tx *gorm.DB) (*User, error) {
	user := User{}
	if err := tx.Where("id = ?", userModel.ID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userModel *User) Select(tx *gorm.DB) (*User, error) {
	user := User{}
	if err := tx.Where("openid = ?", userModel.Openid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userModel *User) Insert(tx *gorm.DB) error {
	if err := tx.Create(userModel).Error; err != nil {
		return err
	}
	return nil
}

func (userModel *User) ExistOpenID(tx *gorm.DB) bool {
	result := tx.Where("openid = ?", userModel.Openid).First(&User{})
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (User) TableName() string {
	return "blind_box_user"
}
