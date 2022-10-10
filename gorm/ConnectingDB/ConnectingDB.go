package main

import (
	"ORM/tool"
	"fmt"
	"log"
)

func main() {
	dbConfig := tool.GetDBConfig()
	_, err := tool.GetGormEngine(dbConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("数据库连接成功")
}
