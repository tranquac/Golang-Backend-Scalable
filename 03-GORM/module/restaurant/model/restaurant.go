package restaurantmodel

import (
	"03-GORM/common"
	"errors"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"
const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name      string             `json:"name" gorm:"column:name;"`
	Addr      string             `json:"addr" gorm:"column:addr;"`
	Type      RestaurantType     `json:"type" gorm:"column:type;"`
	UserId    int                `json:"-" gorm:"column:user_id;"`
	Logo      *(common.Image)    `json:"logo" gorm:"column:logo;"`
	Cover     *common.Images     `json:"cover" gorm:"column:cover;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikeCount int                `json:"liked_count" gorm:"column:liked_count;"`
}

type RestaurantUpdate struct {
	common.SQLModel
	Name  *string         `json:"name" gorm:"column:name;"`
	Addr  *string         `json:"addr" gorm:"column:addr;"`
	Logo  *(common.Image) `json:"logo" gorm:"column:logo"`
	Cover *common.Images  `json:"cover" gorm:"column:cover;"`
}

type RestaurantCreate struct {
	common.SQLModel
	Name   string         `json:"name" gorm:"column:name;"`
	Addr   string         `json:"addr" gorm:"column:addr;"`
	UserId int            `json:"-" gorm:"column:user_id;"`
	Logo   *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover  *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantUpdate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
