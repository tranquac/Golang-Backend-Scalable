package ginrestaurant

import (
	"03-GORM/common"
	restaurantbiz "03-GORM/module/restaurant/biz"
	restaurantstorage "03-GORM/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Delete(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStorage(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)
		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		} else {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		}
	}
}
