package userbiz

import (
	"03-GORM/common"
	usermodel "03-GORM/module/user/model"
	"context"
	"fmt"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, more ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		//if user.Status == 0 {
		//	return error user has been disable
		//}
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	fmt.Println(data.Password)
	fmt.Println(salt)
	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	data.Status = 1

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
