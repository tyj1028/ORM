package main

import (
	"ORM/tool"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
	dbConfig := tool.GetDBConfig()
	_, err := tool.GetXormEngine(dbConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("数据库连接成功")
}
