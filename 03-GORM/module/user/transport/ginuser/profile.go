package ginuser

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
