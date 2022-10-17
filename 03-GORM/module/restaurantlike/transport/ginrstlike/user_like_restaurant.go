package ginrstlike

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	restaurantstorage "03-GORM/module/restaurant/storage"
	rstlikebiz "03-GORM/module/restaurantlike/biz"
	restaurantlikemodel "03-GORM/module/restaurantlike/model"
	restaurantlikestorage "03-GORM/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /v1/restaurants/:id/like

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMaiDBConnection())
		incStore := restaurantstorage.NewSQLStorage(appCtx.GetMaiDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
