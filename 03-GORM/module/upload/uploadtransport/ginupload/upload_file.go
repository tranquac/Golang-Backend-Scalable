package ginupload

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	"03-GORM/module/upload/uploadbusiness"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
)

func Upload(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fileHeader, err := c.FormFile("file")
		//
		//if err != nil {
		//	panic(err)
		//}
		//
		//if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		//	panic(err)
		//}
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
		//	Id:        0,
		//	Url:       "http://localhost:8088/static/" + fileHeader.Filename,
		//	Width:     0,
		//	Height:    0,
		//	CloudName: "local",
		//	Extension: "png",
		//}))
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbusiness.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
