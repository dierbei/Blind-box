package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/dierbei/blind-box/global"
	"github.com/dierbei/blind-box/internal/middleware"
	"github.com/dierbei/blind-box/internal/model"
	"github.com/dierbei/blind-box/pkg/dto"
	"github.com/dierbei/blind-box/pkg/wx"
)

type UserController struct {
}

func UserRegister(router *gin.RouterGroup) {
	user := UserController{}

	userGroup := router.Group("/user")
	userGroup.GET("/login", user.Login)
	userGroup.GET("/myaddlist", user.MyAddList)
	userGroup.GET("/prizelist", user.PrizeList)
	userGroup.GET("/list", middleware.Cors(), user.List)
}

func (handler *UserController) List(ctx *gin.Context) {
	users, err := (&model.User{}).SelectAll(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, 500, err)
		return
	}

	dtoList := make([]dto.UserOutput, 0)
	for _, user := range users {
		userDto := dto.UserOutput{
			ID:        user.ID,
			CreatedAt: model.FormatTime(user.BaseModel.CreatedAt),
			OpenID:    user.Openid,
		}
		dtoList = append(dtoList, userDto)
	}

	middleware.ResponseSuccess(ctx, dtoList)
}

func (handler *UserController) MyAddList(ctx *gin.Context) {
	userID := ctx.DefaultQuery("userid", "0")
	userIDInt, _ := strconv.Atoi(userID)

	if userIDInt == 0 {
		middleware.ResponseError(ctx, 500, errors.New("用户信息不正确"))
		return
	}

	peoples, err := (&model.People{UserID: int32(userIDInt)}).SelectAddListByUserID(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, 500, err)
		return
	}

	dtoList := make([]dto.UserAddListOutput, 0)
	for _, people := range peoples {
		dtoInfo := dto.UserAddListOutput{
			CreatedAt:   model.FormatTime(people.CreatedAt),
			UserID:      people.UserID,
			WxNumber:    people.WxNumber,
			Description: people.Description,
			Local:       people.Local,
		}
		dtoList = append(dtoList, dtoInfo)
	}

	middleware.ResponseSuccess(ctx, dtoList)
}

func (handler *UserController) PrizeList(ctx *gin.Context) {
	userID := ctx.DefaultQuery("userid", "0")
	userIDInt, _ := strconv.Atoi(userID)

	prizes, err := (&model.UserPrize{UserID: int32(userIDInt)}).SelectList(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, 500, err)
		return
	}

	peopleList := make([]*model.People, 0)
	for _, prize := range prizes {
		people, err := (&model.People{BaseModel: model.BaseModel{ID: prize.PrizeID}}).Select(global.MySQLTx)
		if err != nil {
			middleware.ResponseError(ctx, 500, err)
			return
		}

		// 查询people的自拍
		images, _ := (&model.Image{PeopleID: people.ID}).SelectByPeopleID(global.MySQLTx)
		url := make([]string, 0)
		for _, image := range images {
			url = append(url, image.Url)
		}
		people.Images = url

		peopleList = append(peopleList, people)
	}

	dtoList := make([]dto.UserAddListOutput, 0)
	for _, people := range peopleList {
		dtoInfo := dto.UserAddListOutput{
			CreatedAt:   model.FormatTime(people.CreatedAt),
			UserID:      people.UserID,
			WxNumber:    people.WxNumber,
			Description: people.Description,
			Local:       people.Local,
		}
		dtoList = append(dtoList, dtoInfo)
	}

	middleware.ResponseSuccess(ctx, dtoList)
}

func (handler *UserController) Login(ctx *gin.Context) {
	code := ctx.Query("code")
	userSessionInfo, err := wx.WxLogin(code)
	if err != nil {
		middleware.ResponseError(ctx, 500, err)
		return
	}

	user := model.User{Openid: userSessionInfo.OpenID}
	if exist := user.ExistOpenID(global.MySQLTx); !exist {
		if err := user.Insert(global.MySQLTx); err != nil {
			middleware.ResponseError(ctx, 500, err)
			return
		}
	}

	userInfo, err := user.Select(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, 500, err)
		return
	}

	middleware.ResponseSuccess(ctx, dto.UserLoginOutput{
		ID:         userInfo.ID,
		SessionKey: userSessionInfo.SessionKey,
		ExpireIn:   userSessionInfo.ExpireIn,
		OpenID:     userSessionInfo.OpenID,
	})
}
