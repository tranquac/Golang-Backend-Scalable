package restaurantstorage

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
)

func (s *sqlStorage) Get(
	context context.Context,
	id int,
	data *restaurantmodel.Restaurant,
) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id=?", id).
		First(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
