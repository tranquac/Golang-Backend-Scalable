package main

import (
	"03-GORM/component/appctx"
	"03-GORM/middleware"
	"03-GORM/module/restaurant/transport/ginrestaurant"
	"03-GORM/module/restaurantlike/transport/ginrstlike"
	"03-GORM/module/upload/uploadtransport/ginupload"
	"03-GORM/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func SetupMainRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	// upload file
	v1.POST("/upload", ginupload.Upload(appContext))
	// register user
	v1.POST("/register", ginuser.Register(appContext))
	// authenticate
	v1.POST("/authenticate", ginuser.Login(appContext))
	// get profile
	v1.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))
	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appContext))
	//add restaurants
	restaurants.POST("", ginrestaurant.Create(appContext))
	//edit restaurants by ID
	restaurants.PATCH("/:id", ginrestaurant.Edit(appContext.GetMaiDBConnection()))
	//delete restaurants by ID
	restaurants.DELETE("/:id", ginrestaurant.Delete(appContext.GetMaiDBConnection()))
	//get restaurants by id
	restaurants.GET("/:id", ginrestaurant.Get(appContext.GetMaiDBConnection()))
	//get restaurants limit and paging
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	// post like restaurant
	restaurants.POST("/:id/like", ginrstlike.UserLikeRestaurant(appContext))
	// post dislike restaurant
	restaurants.DELETE("/:id/dislike", ginrstlike.UserDisLikeRestaurant(appContext))
	// list user like restaurant GET /v1/restaurants/:id/liked-users
	restaurants.GET("/:id/liked-users", ginrstlike.ListUsers(appContext))

}
