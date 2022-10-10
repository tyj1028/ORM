package main

import (
	"ORM/model"
	"ORM/tool"
	"fmt"
	"log"
)

func main() {
	//查询“c001”课程比“c002”课程成绩高的所有学生的学号；
	dbConfig := tool.GetDBConfig()
	xormEngine, err := tool.GetXormEngine(dbConfig)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer xormEngine.Close()
	xormEngine.ShowSQL(true)

	//1.查询 student 表的性别为男的所有信息
	students := make([]model.Student, 0)
	xormEngine.Where("ssex ='男'").Find(&students)
	fmt.Println(students)

}
