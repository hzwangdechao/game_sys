package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("database.driverName")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect database,err:" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
