package dto

import "github.com/dierbei/blind-box/internal/model"

type PeopleOutput struct {
	ID          int32            `json:"id"`
	CreatedAt   model.FormatTime `json:"created_at"`
	UpdatedAt   model.FormatTime `json:"updated_at"`
	UserID      int32            `json:"user_id"`
	WxNumber    string           `json:"wx_number"`
	Description string           `json:"description"`
	Local       string           `json:"local"`
	Sex         int              `json:"sex"`
	Url         []string         `json:"url"`
}

type PeopleImageOutput struct {
	Url string `json:"url"`
}
