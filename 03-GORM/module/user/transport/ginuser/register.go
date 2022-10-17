package ginuser

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	"03-GORM/component/hasher"
	userbiz "03-GORM/module/user/biz"
	usermodel "03-GORM/module/user/model"
	userstore "03-GORM/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMD5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
