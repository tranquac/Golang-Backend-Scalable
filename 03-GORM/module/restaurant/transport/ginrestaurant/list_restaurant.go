package ginrestaurant

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	restaurantbiz "03-GORM/module/restaurant/biz"
	restaurantmodel "03-GORM/module/restaurant/model"
	restaurantrepo "03-GORM/module/restaurant/repository"
	restaurantstorage "03-GORM/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStorage(db)
		//likeStore := restaurantlikestorage.NewSqlStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
