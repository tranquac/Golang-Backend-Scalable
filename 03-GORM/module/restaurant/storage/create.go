package restaurantstorage

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
)

func (s *sqlStorage) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
