package v1

import (
	"github.com/dierbei/blind-box/pkg/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/dierbei/blind-box/global"
	"github.com/dierbei/blind-box/internal/middleware"
	"github.com/dierbei/blind-box/internal/model"
	"github.com/dierbei/blind-box/pkg/forms"
)

type PeopleController struct {
}

func ManRegister(router *gin.RouterGroup) {
	people := PeopleController{}

	peopleGroup := router.Group("/people")
	peopleGroup.GET("/add", people.AddPeople)
	peopleGroup.GET("/random", people.RandomPeople)
	peopleGroup.GET("/list", people.List)
}

func (handler *PeopleController) List(ctx *gin.Context) {
	peoples, err := (&model.People{}).SelectAll(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	peopleDtoList := make([]dto.PeopleOutput, 0)
	for _, people := range peoples {
		peopleDto := dto.PeopleOutput{
			ID:          people.BaseModel.ID,
			CreatedAt:   model.FormatTime(people.BaseModel.CreatedAt),
			UpdatedAt:   model.FormatTime(people.BaseModel.UpdatedAt),
			UserID:      people.UserID,
			WxNumber:    people.WxNumber,
			Description: people.Description,
			Local:       people.Local,
		}
		peopleDtoList = append(peopleDtoList, peopleDto)
	}

	middleware.ResponseSuccess(ctx, peopleDtoList)
}

func (handler *PeopleController) RandomPeople(ctx *gin.Context) {
	userID := ctx.DefaultQuery("userid", "0")
	userIDInt, _ := strconv.Atoi(userID)
	sex := ctx.DefaultQuery("sex", "0")
	sexInt, _ := strconv.Atoi(sex)

	if userIDInt == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("用户信息不正确"))
		return
	}

	people, err := (&model.People{UserID: int32(userIDInt), Sex: sexInt}).Random(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	middleware.ResponseSuccess(ctx, dto.PeopleOutput{
		CreatedAt:   model.FormatTime(people.CreatedAt),
		UpdatedAt:   model.FormatTime(people.UpdatedAt),
		UserID:      people.UserID,
		WxNumber:    people.WxNumber,
		Description: people.Description,
		Local:       people.Local,
	})
	return
}

func (handler *PeopleController) AddPeople(ctx *gin.Context) {
	userID := ctx.DefaultQuery("userid", "0")
	userIDInt, _ := strconv.Atoi(userID)
	sex := ctx.DefaultQuery("sex", "0")
	sexInt, _ := strconv.Atoi(sex)

	if userIDInt == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("用户信息错误"))
		return
	}

	params := &forms.PeopleAddForm{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	if exist := (&model.People{WxNumber: params.WxNumber}).SelectByWxNumber(global.MySQLTx); exist {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("微信号已存在"))
		return
	}

	err := (&model.People{
		Images:      params.Images,
		WxNumber:    params.WxNumber,
		Description: params.Description,
		Local:       params.Local,
		UserID:      int32(userIDInt),
		Sex:         sexInt,
	}).Insert(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	middleware.ResponseSuccess(ctx, "")
	return
}
