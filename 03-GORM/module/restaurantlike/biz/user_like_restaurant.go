package rstlikebiz

import (
	restaurantlikemodel "03-GORM/module/restaurantlike/model"
	"context"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncLikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountResStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncLikedCountResStore,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, incStore: incStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
		log.Println(err)
	}
	return nil
}
