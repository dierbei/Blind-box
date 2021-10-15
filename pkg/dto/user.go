package dto

import (
	"github.com/dierbei/blind-box/internal/model"
)

type UserLoginOutput struct {
	ID         int32  `json:"id"`
	SessionKey string `json:"session_key"`
	ExpireIn   int    `json:"expires_in"`
	OpenID     string `json:"openid"`
}

type UserOutput struct {
	ID        int32            `json:"id"`
	CreatedAt model.FormatTime `json:"created_at"`
	OpenID    string           `json:"open_id"`
}

type UserAddListOutput struct {
	ID          int32            `json:"id"`
	CreatedAt   model.FormatTime `json:"created_at"`
	UpdatedAt   model.FormatTime `json:"updated_at"`
	UserID      int32            `json:"user_id"`
	WxNumber    string           `json:"wx_number"`
	Description string           `json:"description"`
	Local       string           `json:"local"`
	Images      []string         `json:"images"`
}

type UserPrizeListOutput struct {
	model.BaseModel
	UserID      int32  `json:"user_id"`
	WxNumber    string `json:"wx_number"`
	Description string `json:"description"`
	Local       string `json:"local"`
}
