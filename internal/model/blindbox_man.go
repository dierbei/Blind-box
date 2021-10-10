package model

type Man struct {
	BaseModel
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

func (Man) TableName() string {
	return "blind_box_man"
}
