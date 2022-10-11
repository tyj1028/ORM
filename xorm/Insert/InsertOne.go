package main

import (
	"ORM/model"
	"ORM/tool"
	"fmt"
	"log"
)

func main() {
	dbConfig := tool.GetDBConfig()
	xormEngine, err := tool.GetXormEngine(dbConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	student := model.Student{
		Sno:   "s014",
		Sname: "钟杰",
		Sage:  25,
		Ssex:  "男",
	}
	insert, err := xormEngine.Insert(&student)
	fmt.Println(insert)
}
