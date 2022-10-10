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

	//1.查询 student 表的性别为男的所有信息
	//students := make([]model.Student, 0)
	//gormEngine.Where("ssex ='男'").Find(&students)
	//fmt.Println(students)

	//1.查询 student 表的性别为男的所有信息
	students := make([]model.Student, 0)
	gormEngine.Where(model.Student{Ssex: "男"}).Find(&students)
	fmt.Println(students)

}
