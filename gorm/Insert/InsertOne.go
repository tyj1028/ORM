package main

import (
	"ORM/model"
	"ORM/tool"
	"fmt"
	"log"
)

func main() {
	dbConfig := tool.GetDBConfig()
	gormEngine, err := tool.GetGormEngine(dbConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	//插入一条信息
	student := model.Student{
		Sno:   "s012",
		Sname: "张杰",
		Sage:  25,
		Ssex:  "男",
	}
	tx := gormEngine.Create(&student)
	fmt.Printf("%d条记录被更新", tx.RowsAffected)
}
