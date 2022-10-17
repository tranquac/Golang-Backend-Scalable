package restaurantbiz

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
)

type EditRestaurantStore interface {
	Edit(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type editRestaurantBiz struct {
	store EditRestaurantStore
}

func NewEditRestaurantBiz(store EditRestaurantStore) *editRestaurantBiz {
	return &editRestaurantBiz{store: store}
}

func (biz *editRestaurantBiz) EditRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := biz.store.Edit(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, nil)
	}
	return nil
}
