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

func (peopleModel *People) Delete(tx *gorm.DB) error {
	if err := tx.Table(peopleModel.TableName()).Where("id = ?", peopleModel.ID).Delete(&People{}).Error; err != nil {
		return err
	}
	return nil
}

func (peopleModel *People) SelectAll(tx *gorm.DB) ([]People, error) {
	peoples := make([]People, 0)
	if err := tx.Table(peopleModel.TableName()).Find(&peoples).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(peoples); i++ {
		images, _ := (&Image{PeopleID: peoples[i].ID}).SelectByPeopleID(tx)
		urlList := make([]string, 0)
		for _, image := range images {
			urlList = append(urlList, image.Url)
		}
		peoples[i].Images = urlList
	}

	//for _, people := range peoples {
	//	images, _ := (&Image{}).SelectByPeopleID(tx)
	//	urlList := make([]string, 0)
	//	for _, image := range images {
	//		urlList = append(urlList, image.Url)
	//	}
	//}

	return peoples, nil
}

func (peopleModel *People) Random(tx *gorm.DB) (*People, error) {
	people := People{}

	tx.Raw("SELECT * FROM blind_box_people WHERE sex = ? ORDER BY RAND() limit 1", peopleModel.Sex).Scan(&people)
	if people.ID == 0 {
		return nil, errors.New("没有查询到任何信息")
	}

	// 查询此奖品和此用户是否有关联关系，如果有直接返回
	//if exist := (&UserPrize{UserID: peopleModel.UserID, PrizeID: people.ID}).Exist(tx); exist {
	//	// 查询奖品的自拍
	//	images, err := (&Image{PeopleID: peopleModel.ID}).SelectByPeopleID(tx)
	//	if err != nil {
	//		return &people, nil
	//	}
	//	url := make([]string, 0)
	//	for _, image := range images {
	//		url = append(url, image.Url)
	//	}
	//	people.Images = url
	//	return &people, nil
	//}

	// 创建奖品和用户的关联关系
	//if err := (&UserPrize{UserID: peopleModel.UserID, PrizeID: people.ID}).Insert(tx); err != nil {
	//	return nil, err
	//}

	return &people, nil
}

func (peopleModel *People) SelectByWxNumber(tx *gorm.DB) bool {
	people := People{}
	result := tx.Table(peopleModel.TableName()).Where("wx_number = ?", peopleModel.WxNumber).Find(&people)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (peopleModel *People) Select(tx *gorm.DB) (*People, error) {
	people := People{}
	result := tx.Table(peopleModel.TableName()).Where("id = ?", peopleModel.ID).Order("created_at desc").Find(&people)
	if result.Error != nil {
		return nil, result.Error
	}
	return &people, nil
}

func (peopleModel *People) Insert(tx *gorm.DB) error {
	newTx := tx.Begin()

	peopleModel.Images = nil
	if err := newTx.Table(peopleModel.TableName()).Create(peopleModel).Error; err != nil {
		return err
	}

	//fileSlice := strings.Split(fileList, ",")

	//if len(fileSlice) > 9 {
	//	return errors.New("最多上传9张图片")
	//}
	//
	////images := make([]Image, 0)
	//if len(fileSlice) != 0 {
	//	for _, url := range fileSlice {
	//		image := Image{
	//			PeopleID: peopleModel.ID,
	//			Url:      url,
	//		}
	//		if err := image.Insert(tx); err != nil {
	//			tx.Rollback()
	//			return err
	//		}
	//	}
	//}

	//if len(images) > 0 {
	//	if err := (&Image{}).Insert(newTx, images); err != nil {
	//		newTx.Rollback()
	//		return err
	//	}
	//}

	newTx.Commit()
	return nil
}

func (peopleModel *People) SelectAddListByUserID(tx *gorm.DB) ([]People, error) {
	peoples := make([]People, 0)
	result := tx.Table(peopleModel.TableName()).Where("user_id = ?", peopleModel.UserID).Order("created_at desc").Find(&peoples)
	if result.Error != nil {
		return nil, result.Error
	}

	//for i := 0; i < len(peoples); i++ {
	//	// 查询people的自拍
	//	images, _ := (&Image{PeopleID: peoples[i].ID}).SelectByPeopleID(tx)
	//	url := make([]string, 0)
	//	for _, image := range images {
	//		url = append(url, image.Url)
	//	}
	//	peoples[i].Images = url
	//}

	return peoples, result.Error
}

func (peopleModel *People) PageList(tx *gorm.DB, params *forms.ManListPageInput) ([]People, int64, error) {
	list := make([]People, 0)
	var count int64
	offset := (params.Page - 1) * params.PageSize

	result := tx.Limit(params.PageSize).Offset(offset).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, 0, result.Error
	}

	result = tx.Table(peopleModel.TableName()).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return list, count, nil
}

func (peopleModel *People) SelectByUserID(tx *gorm.DB, id int64) ([]People, error) {
	list := make([]People, 0)

	result := tx.Where("user_id = ?", id).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	return list, nil
}

func (peopleModel *People) TableName() string {
	return "blind_box_people"
}
