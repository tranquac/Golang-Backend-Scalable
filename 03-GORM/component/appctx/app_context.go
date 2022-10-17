package appctx

import (
	"03-GORM/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
