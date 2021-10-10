package model

type Woman struct {
	BaseModel
	UserID   int32  `json:"user_id"`
	Username string `json:"username"`
}

func (Woman) TableName() string {
	return "blind_box_woman"
}
