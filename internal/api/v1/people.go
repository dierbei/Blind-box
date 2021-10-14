package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/dierbei/blind-box/global"
	"github.com/dierbei/blind-box/internal/middleware"
	"github.com/dierbei/blind-box/internal/model"
	"github.com/dierbei/blind-box/pkg/ali_oss"
	"github.com/dierbei/blind-box/pkg/dto"
	"github.com/dierbei/blind-box/pkg/forms"
)

type PeopleController struct {
}

func ManRegister(router *gin.RouterGroup) {
	people := PeopleController{}

	peopleGroup := router.Group("/people")
	peopleGroup.GET("/add", people.AddPeople)
	peopleGroup.GET("/random", people.RandomPeople)
	peopleGroup.GET("/list", middleware.Cors(), people.List)
	peopleGroup.POST("/upload", people.Upload)
}

func (handler *PeopleController) Upload(ctx *gin.Context) {
	filePath := ali_oss.UploadFile(ctx)
	middleware.ResponseSuccess(ctx, dto.PeopleImageOutput{Url: filePath})
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
			Url:         people.Images,
			Sex:         people.Sex,
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

	// 判断用户是否有权限抽取纸条
	user, err := (&model.User{BaseModel: model.BaseModel{ID: int32(userIDInt)}}).SelectByID(global.MySQLTx)
	if err != nil || user.Status == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("需要先投递一次纸条哦~"))
		return
	}

	people, err := (&model.People{UserID: int32(userIDInt), Sex: sexInt}).Random(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	middleware.ResponseSuccess(ctx, dto.PeopleOutput{
		ID:          people.ID,
		CreatedAt:   model.FormatTime(people.CreatedAt),
		UpdatedAt:   model.FormatTime(people.UpdatedAt),
		UserID:      people.UserID,
		WxNumber:    people.WxNumber,
		Description: people.Description,
		Local:       people.Local,
		Sex:         people.Sex,
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

	if err = (&model.User{}).UpdateStatus(global.MySQLTx); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	middleware.ResponseSuccess(ctx, nil)
	return
}
