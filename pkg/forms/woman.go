package forms

type WomanListPageInput struct {
	PageSize int `form:"page_size" json:"page_size" comment:"每页记录数" validate:"" example:"10"`
	Page     int `form:"page" json:"page" comment:"页数" validate:"required" example:"1"`
}
