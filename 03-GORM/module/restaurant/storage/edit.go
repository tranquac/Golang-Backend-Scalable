package restaurantstorage

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
)

func (s *sqlStorage) Edit(
	context context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := s.db.Table(restaurantmodel.RestaurantUpdate{}.TableName()).
		Where("id=?", id).
		Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
