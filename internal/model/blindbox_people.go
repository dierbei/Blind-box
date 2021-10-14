package model

import (
	"github.com/dierbei/blind-box/pkg/forms"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type People struct {
	BaseModel
	UserID      int32    `json:"user_id"`
	WxNumber    string   `json:"wx_number"`
	Description string   `json:"description"`
	Local       string   `json:"local"`
	Sex         int      `json:"sex"`
	Images      []string `json:"images"`
}

func (model *People) SelectAll(tx *gorm.DB) ([]People, error) {
	peoples := make([]People, 0)
	if err := tx.Table("blind_box_people").Find(&peoples).Error; err != nil {
		return nil, err
	}
	return peoples, nil
}

func (model *People) Random(tx *gorm.DB) (*People, error) {
	people := People{}

	tx.Raw("SELECT * FROM blind_box_people WHERE sex = ? ORDER BY RAND() limit 1", model.Sex).Scan(&people)
	if people.ID == 0 {
		return nil, errors.New("没有查询到任何信息")
	}

	// 查询此奖品和此用户是否有关联关系，如果有直接返回
	if exist := (&UserPrize{UserID: model.UserID, PrizeID: people.ID}).Exist(tx); exist {
		// 查询奖品的自拍
		images, err := (&Image{PeopleID: model.ID}).SelectByPeopleID(tx)
		if err != nil {
			return &people, nil
		}
		url := make([]string, 0)
		for _, image := range images {
			url = append(url, image.Url)
		}
		people.Images = url
		return &people, nil
	}

	// 创建奖品和用户的关联关系
	if err := (&UserPrize{UserID: model.UserID, PrizeID: people.ID}).Insert(tx); err != nil {
		return nil, err
	}

	return &people, nil
}

func (model *People) SelectByWxNumber(tx *gorm.DB) bool {
	people := People{}
	result := tx.Where("wx_number = ?", model.WxNumber).Find(&people)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (model *People) Select(tx *gorm.DB) (*People, error) {
	people := People{}
	result := tx.Where("id = ?", model.ID).Find(&people)
	if result.Error != nil {
		return nil, result.Error
	}
	return &people, nil
}

func (model *People) Insert(tx *gorm.DB) error {
	newTx := tx.Begin()

	if err := newTx.Create(model).Error; err != nil {
		return err
	}

	if len(model.Images) > 9 {
		return errors.New("最多上传9张图片")
	}

	images := make([]Image, 0)
	if len(model.Images) != 0 {
		for _, url := range model.Images {
			image := Image{
				PeopleID: model.ID,
				Url:      url,
			}
			images = append(images, image)
		}
	}

	if err := (&Image{}).Insert(newTx, images); err != nil {
		newTx.Rollback()
		return err
	}

	newTx.Commit()
	return nil
}

func (model *People) SelectAddListByUserID(tx *gorm.DB) ([]People, error) {
	peoples := make([]People, 0)
	result := tx.Where("user_id = ?", model.UserID).Find(&peoples)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := 0; i < len(peoples); i++ {
		// 查询people的自拍
		images, _ := (&Image{PeopleID: peoples[i].ID}).SelectByPeopleID(tx)
		url := make([]string, 0)
		for _, image := range images {
			url = append(url, image.Url)
		}
		peoples[i].Images = url
	}

	return peoples, result.Error
}

func (model *People) PageList(tx *gorm.DB, params *forms.ManListPageInput) ([]People, int64, error) {
	list := make([]People, 0)
	var count int64
	offset := (params.Page - 1) * params.PageSize

	result := tx.Limit(params.PageSize).Offset(offset).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, 0, result.Error
	}

	result = tx.Table(model.TableName()).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return list, count, nil
}

func (model *People) SelectByUserID(tx *gorm.DB, id int64) ([]People, error) {
	list := make([]People, 0)

	result := tx.Where("user_id = ?", id).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	return list, nil
}

func (model *People) TableName() string {
	return "blind_box_people"
}
