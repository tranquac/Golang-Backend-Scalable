package middleware

import (
	"03-GORM/common"
	"03-GORM/component/appctx"
	"03-GORM/component/tokenprovider/jwt"
	userstore "03-GORM/module/user/store"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(err,
		fmt.Sprintf("wrong auth header"),
		fmt.Sprintf("ErrWrongAuthHeader"))
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to Payload
// 3. From the token payload, we use user_id to find from DB
func RequiredAuth(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMaiDBConnection()
		store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been delete or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
