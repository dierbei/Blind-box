package ali_oss

import (
	"bytes"

	"github.com/gin-gonic/gin"

	"github.com/dierbei/blind-box/global"
)

func UploadFile(ctx *gin.Context) string {
	file, err := ctx.FormFile("file")

	picBuf := make([]byte, file.Size)
	open, err := file.Open()
	_, err = open.Read(picBuf)
	if err != nil {
		return ""
	}
	// 获取存储空间。
	bucket, err := global.AliOSS.Bucket("manhe-hcyj")
	if err != nil {
		return ""
	}

	// 上传Byte数组。
	err = bucket.PutObject(file.Filename, bytes.NewReader([]byte(picBuf)))
	if err != nil {
		return ""
	}

	return global.ServerConfig.AliOSS.BaseUrl + file.Filename
}
