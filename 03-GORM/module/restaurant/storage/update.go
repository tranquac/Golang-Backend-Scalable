package restaurantstorage

import (
	"03-GORM/common"
	restaurantmodel "03-GORM/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStorage) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStorage) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
