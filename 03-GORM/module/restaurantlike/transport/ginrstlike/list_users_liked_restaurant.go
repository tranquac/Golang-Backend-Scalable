package ginrstlike

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	rstlikebiz "03-GORM/module/restaurantlike/biz"
	restaurantlikemodel "03-GORM/module/restaurantlike/model"
	restaurantlikestorage "03-GORM/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Filfill()

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMaiDBConnection())
		biz := rstlikebiz.NewListUsersLikedRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
