package rstlikebiz

import (
	"03-GORM/common"
	restaurantlikemodel "03-GORM/module/restaurantlike/model"
	"context"
	"log"
	"time"
)

type UserDisLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecLikedCountResStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDisLikeRestaurantBiz struct {
	store    UserDisLikeRestaurantStore
	DecStore DecLikedCountResStore
}

func NewUserDisLikeRestaurantBiz(store UserDisLikeRestaurantStore, DecStore DecLikedCountResStore) *userDisLikeRestaurantBiz {
	return &userDisLikeRestaurantBiz{store: store, DecStore: DecStore}
}

func (biz *userDisLikeRestaurantBiz) DisLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.DecStore.DecreaseLikeCount(ctx, restaurantId); err != nil {
			log.Println(err)

			for i := 1; i <= 3; i++ {
				err := biz.DecStore.DecreaseLikeCount(ctx, restaurantId)
				if err == nil {
					log.Println(err)
					break
				}
				time.Sleep(3 * time.Second)
			}
		}
	}()

	return nil
}
