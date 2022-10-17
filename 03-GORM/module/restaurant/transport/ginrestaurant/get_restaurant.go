package ginrestaurant

import (
	"03-GORM/common"
	restaurantbiz "03-GORM/module/restaurant/biz"
	restaurantmodel "03-GORM/module/restaurant/model"
	restaurantstorage "03-GORM/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Get(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.Restaurant

		store := restaurantstorage.NewSQLStorage(db)
		biz := restaurantbiz.NewGetRestaurantBiz(store)
		if err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
