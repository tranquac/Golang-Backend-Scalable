package ginuser

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	"03-GORM/component/hasher"
	"03-GORM/component/tokenprovider/jwt"
	userbiz "03-GORM/module/user/biz"
	usermodel "03-GORM/module/user/model"
	userstore "03-GORM/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMaiDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMD5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
