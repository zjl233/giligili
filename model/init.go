package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

// OSS bucket 单例
var BKT *oss.Bucket

func OSS(endpoint, accessKeyID, accessKeySecret, bucketName string) {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}

	BKT = bucket
}
