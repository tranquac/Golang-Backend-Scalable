package restaurantlikestorage

import (
	"03-GORM/common"
	restaurantlikemodel "03-GORM/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		panic(common.ErrDB(err))
	}
	return nil
}
