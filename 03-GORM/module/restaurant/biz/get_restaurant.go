package restaurantbiz

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
)

type GetRestaurantStore interface {
	Get(context context.Context, id int, data *restaurantmodel.Restaurant) error
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int, data *restaurantmodel.Restaurant) error {
	if err := biz.store.Get(ctx, id, data); err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, nil)
	}
	return nil
}
