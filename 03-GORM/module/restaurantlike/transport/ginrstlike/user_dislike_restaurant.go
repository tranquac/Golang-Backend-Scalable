package ginrstlike

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	restaurantstorage "03-GORM/module/restaurant/storage"
	rstlikebiz "03-GORM/module/restaurantlike/biz"
	restaurantlikestorage "03-GORM/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /v1/restaurants/:id/dislike

func UserDisLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMaiDBConnection())
		DecStore := restaurantstorage.NewSQLStorage(appCtx.GetMaiDBConnection())
		biz := rstlikebiz.NewUserDisLikeRestaurantBiz(store, DecStore)

		if err := biz.DisLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
