package ginrestaurant

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	restaurantbiz "03-GORM/module/restaurant/biz"
	restaurantmodel "03-GORM/module/restaurant/model"
	restaurantstorage "03-GORM/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		data.Status = 1

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStorage(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
