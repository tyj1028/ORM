package tool

import "ORM/config"

func GetDBConfig() config.DbConfig {
	config, err := config.ParseConfig("config/config.json")
	configDB := config.Database
	if err != nil {
		panic(err.Error())
	}
	return configDB
}
