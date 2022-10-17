package restaurantrepo

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

//type LikeRestaurantStore interface {
//	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
//}

type ListRestaurantRepo struct {
	store ListRestaurantStore
	//likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *ListRestaurantRepo {
	return &ListRestaurantRepo{store: store}
}

func (biz *ListRestaurantRepo) ListRestaurant(context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))
	//for i := range ids {
	//	ids[i] = result[i].Id
	//}
	//likeMap, err := biz.likeStore.GetRestaurantLikes(context, ids)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//for i, item := range result {
	//	result[i].LikeCount = likeMap[item.Id]
	//}

	return result, nil
}
