package restaurantstorage

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
)

func (s *sqlStorage) Delete(
	context context.Context,
	id int,
) error {

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id=?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
