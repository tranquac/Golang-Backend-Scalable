package rstlikebiz

import (
	"03-GORM/common"
	restaurantlikemodel "03-GORM/module/restaurantlike/model"
	"context"
)

type ListUsersLikedRestaurant interface {
	GetUsersLikedRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUsersLikedRestaurantBiz struct {
	store ListUsersLikedRestaurant
}

func NewListUsersLikedRestaurantBiz(store ListUsersLikedRestaurant) *listUsersLikedRestaurantBiz {
	return &listUsersLikedRestaurantBiz{
		store: store,
	}
}

func (biz *listUsersLikedRestaurantBiz) ListUsers(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikedRestaurant(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
