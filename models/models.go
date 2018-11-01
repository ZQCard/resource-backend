package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"resource-backend/pkg/config"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key"json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func init()  {
	var (
		err error
		dbType, dbName, dbUser, dbPassword, dbHost string
	)

	dbType = config.DatabaseSetting.Type
	dbName = config.DatabaseSetting.Name
	dbUser = config.DatabaseSetting.User
	dbPassword = config.DatabaseSetting.Password
	dbHost = config.DatabaseSetting.Host

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
	))

	if err != nil {
		log.Fatalf("数据库连接失败, : %s", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxIdleConns(100)
}

func CloseDB()  {
	defer db.Close()
}


