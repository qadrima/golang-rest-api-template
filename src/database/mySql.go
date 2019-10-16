package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

// InitMysql is
func InitMysql() {
	//open a mysql connection
	var err error
	db, err = gorm.Open("mysql", viper.GetString("database.mysql.user")+":"+viper.GetString("database.mysql.password")+"@tcp("+viper.GetString("database.mysql.host")+":"+viper.GetString("database.mysql.port")+")/"+viper.GetString("database.mysql.dbname")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

// GetDB is
func MysqlDB() *gorm.DB {
	return db
}
