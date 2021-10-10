package model

type User struct {
	BaseModel
	Openid string `json:"openid"`
}

func (User) TableName() string {
	return "blind_box_user"
}
