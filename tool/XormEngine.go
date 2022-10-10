package tool

import (
	"ORM/config"
	"xorm.io/xorm"
)

func GetXormEngine(configDB config.DbConfig) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(configDB.Driver, configDB.User+":"+configDB.Password+"@/"+configDB.DbName+"?charset="+configDB.Charset)
	return engine, err
}
