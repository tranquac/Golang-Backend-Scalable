package userbiz

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	"03-GORM/component/tokenprovider"
	usermodel "03-GORM/module/user/model"
	"context"
	"fmt"
)

type LoginStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, more ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Privider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)
func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	fmt.Println(user)
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)
	fmt.Println(passHashed)
	if user.Password != passHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
