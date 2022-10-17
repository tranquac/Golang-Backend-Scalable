package main

import (
	"03-GORM/component/appctx"
	"03-GORM/component/uploadprovider"
	"03-GORM/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"` //tag
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dns := "food_delivery:quac@tcp(103.81.86.132:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()
	s3BucketName := "quac-s3"
	s3Region := "us-west-1"
	s3APIkey := "AKIAR6ZKZUHKVNGS5JWI"
	s3SecretKey := "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"
	s3Domain := "https://d1c2zbq9fgzt87.cloudfront.net"
	secretKey := "tranvanquac_secret"
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIkey, s3SecretKey, s3Domain)
	appContext := appctx.NewAppContext(db, s3Provider, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	v1 := r.Group("v1")
	r.Static("/static", "./static")
	SetupMainRoute(appContext, v1)
	SetupAdminRoute(appContext, v1)
	r.Run("0.0.0.0:8088")
}
