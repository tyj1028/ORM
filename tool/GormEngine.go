package tool

import (
	"ORM/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetGormEngine(configDB config.DbConfig) (*gorm.DB, error) {
	//dsn:="root:123456@tcp(127.0.0.1:3306)/orm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := configDB.User + ":" + configDB.Password + "@tcp(" + configDB.Host + ":" + configDB.Post + ")/" + configDB.DbName + "?charset=" + configDB.Charset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
