package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"

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
	peopleGroup.GET("/add", middleware.Logging(), people.AddPeople)
	peopleGroup.GET("/random", middleware.Logging(), people.RandomPeople)
	peopleGroup.GET("/list", middleware.Logging(), middleware.Cors(), people.List)
	peopleGroup.POST("/upload", people.Upload)

	peopleGroup.GET("/delete", middleware.Logging(), people.Delete)
}

func (handler *PeopleController) Delete(ctx *gin.Context) {
	userID := ctx.DefaultQuery("userid", "0")
	userIDInt, _ := strconv.Atoi(userID)
	peopleID := ctx.DefaultQuery("pid", "0")
	peopleIDInt, _ := strconv.Atoi(peopleID)

	if userIDInt == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("请先登录"))
		return
	}
	if peopleIDInt == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("无效参数"))
		return
	}

	people, err := (&model.People{BaseModel: model.BaseModel{ID: int32(peopleIDInt)}}).Select(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	if int32(userIDInt) != people.UserID {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("无效参数"))
		return
	}

	if err := (&model.People{BaseModel: model.BaseModel{ID: int32(peopleIDInt)}}).Delete(global.MySQLTx); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("删除失败"))
		return
	}

	middleware.ResponseSuccess(ctx, nil)
}

func (handler *PeopleController) Upload(ctx *gin.Context) {
	//peopleid := ctx.DefaultQuery("pid", "0")
	//fmt.Println(peopleid)
	imageInput := forms.ImageInput{}
	if err := ctx.ShouldBind(&imageInput); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("无效参数"))
		return
	}

	//fmt.Println(imageInput.Pid, "-----------------------------------------------------------------")

	peopleidInt, _ := strconv.Atoi(imageInput.Pid)

	if peopleidInt == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("无效参数"))
		return
	}

	//fileList := ctx.DefaultQuery("fileList", "")

	//var newFileList string
	//if fileList != "" {
	//	newFileList = fileList[1:]
	//}

	//fileSlice := strings.Split(newFileList, ",")

	//if len(fileSlice) > 9 {
	//	middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("最多上传9张图片"))
	//	return
	//}

	url := ali_oss.UploadFile(ctx)

	if err := (&model.Image{PeopleID: int32(peopleidInt), Url: url}).Insert(global.MySQLTx); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("上传图片失败"))
		return
	}

	//images := make([]Image, 0)
	//if len(fileSlice) != 0 {
	//	for _, url := range fileSlice {
	//		image := model.Image{
	//			PeopleID: int32(peopleidInt),
	//			Url:      url,
	//		}
	//		if err := image.Insert(global.MySQLTx); err != nil {
	//			middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("上传图片失败"))
	//			return
	//		}
	//	}
	//}

	//if len(images) > 0 {
	//	if err := (&Image{}).Insert(newTx, images); err != nil {
	//		newTx.Rollback()
	//		return err
	//	}
	//}

	ali_oss.UploadFile(ctx)
	ctx.JSON(http.StatusOK, nil)
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

	if exist := (&model.UserPrize{UserID: int32(userIDInt), PrizeID: people.ID}).Exist(global.MySQLTx); !exist {
		if err := (&model.UserPrize{UserID: int32(userIDInt), PrizeID: people.ID}).Insert(global.MySQLTx); err != nil {
			middleware.ResponseError(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	images, err := (&model.Image{PeopleID: people.ID}).SelectByPeopleID(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	urls := make([]string, 0)
	for _, image := range images {
		urls = append(urls, image.Url)
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
		Url:         urls,
	})
	return
}

func (handler *PeopleController) AddPeople(ctx *gin.Context) {
	userID := ctx.DefaultQuery("userid", "0")
	userIDInt, _ := strconv.Atoi(userID)
	sex := ctx.DefaultQuery("sex", "0")
	sexInt, _ := strconv.Atoi(sex)

	//fileList := ctx.DefaultQuery("fileList", "")
	//var newFileList string
	//if fileList != "" {
	//	newFileList = fileList[1:]
	//}

	// 不正确直接返回
	if userIDInt == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("用户未登录"))
		return
	}

	// 判断用户是否存在
	user, _ := (&model.User{BaseModel: model.BaseModel{ID: int32(userIDInt)}}).SelectByID(global.MySQLTx)
	if user == nil || user.ID == 0 {
		middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("用户不存在"))
		return
	}

	// 数据校验
	params := &forms.PeopleAddForm{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	// 查看微信号是否存在
	//if exist := (&model.People{WxNumber: params.WxNumber}).SelectByWxNumber(global.MySQLTx); exist {
	//	middleware.ResponseError(ctx, http.StatusInternalServerError, errors.New("微信号已存在"))
	//	return
	//}

	people := &model.People{
		//Images:      params.Images,
		WxNumber:    params.WxNumber,
		Description: params.Description,
		Local:       params.Local,
		UserID:      int32(userIDInt),
		Sex:         sexInt,
	}
	err := people.Insert(global.MySQLTx)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	//err := (&model.Image{}).Insert(global.MySQLTx)

	// 更新用户状态
	if err = (&model.User{BaseModel: model.BaseModel{ID: int32(userIDInt)}}).UpdateStatus(global.MySQLTx); err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	middleware.ResponseSuccess(ctx, people)
	return
}
