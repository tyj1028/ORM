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
	//1.查询表 student 姓名为王丽的所有信息
	//var student model.Student
	//gormEngine.Where("sname='王丽'").Find(&student)
	//fmt.Println(student)

	//1.查询表 student 姓名为王丽的所有信息
	//var student model.Student
	//gormEngine.Where("sname=?", "王丽").Find(&student)
	//fmt.Println(student)

	//1.查询表 student 姓名为王丽的所有信息
	//var student model.Student
	//gormEngine.Where("sname=?", "王丽").Find(&student)
	//fmt.Println(student)

	//1.查询表 student 姓名为王丽的所有信息
	var student model.Student
	gormEngine.Where(model.Student{Sname: "王丽"}).Find(&student)
	fmt.Println(student)

	//2.查询表 student 姓名为王丽的学号
	//var student model.Student
	//gormEngine.Select("sno").Where("sname=?", "王丽").Find(&student)
	//fmt.Println(student)

	//2.查询表 student 姓名为王丽的学号
	//var student model.Student
	//gormEngine.Omit("sname,sage,ssex").Where("sname=?", "王丽").Find(&student)
	//fmt.Println(student)

	//3.查询表 student 姓名为王丽的学号和姓名
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//gormEngine.Select("sno,sname").Where("sname =?", name).Find(&student)
	//fmt.Println(student)

	//3.查询表 student 姓名为王丽的学号和姓名
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//gormEngine.Select("sno", "sname").Where("sname =?", name).Find(&student)
	//fmt.Println(student)

}
