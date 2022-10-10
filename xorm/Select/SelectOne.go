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
	defer xormEngine.Close()
	xormEngine.ShowSQL(true)
	//1.查询表 student 姓名为王丽的所有信息
	student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	xormEngine.Where("sname ='王丽'").Get(student)
	fmt.Println(student)

	//1.查询表 student 姓名为王丽的所有信息
	//student := new(model.Student)  //只有一个所以查出只有一个姓名为王丽的人
	//xormEngine.Where("sname =?", "王丽").Get(student)
	//fmt.Println(student)

	//1.查询表 student 姓名为王丽的所有信息
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//xormEngine.Where("sname =?", name).Get(student)
	//fmt.Println(student)

	//1.查询表 student 姓名为王丽的所有信息
	//ERROR
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//xormEngine.Where(model.Student{Sname: name}).Get(student)
	//fmt.Println(student)

	//2.查询表 student 姓名为王丽的学号
	//student := new(model.Student)//只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//xormEngine.Cols("sno").Where("sname =?", name).Get(student)
	//fmt.Println(student)

	//2.查询表 student 姓名为王丽的学号
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//xormEngine.Omit("sname,sage,ssex").Where("sname =?", name).Get(student)
	//fmt.Println(student)

	//2.查询表 student 姓名为王丽的学号
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//xormEngine.Select("sno").Where("sname =?", name).Get(student)
	//fmt.Println(student)

	//3.查询表 student 姓名为王丽的学号和姓名
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//xormEngine.Select("sno,sname").Where("sname =?", name).Get(student)
	//fmt.Println(student)

	//3.查询表 student 姓名为王丽的学号和姓名
	//ERROR
	//student := new(model.Student) //只有一个所以查出只有一个姓名为王丽的人
	//name := "王丽"
	//xormEngine.Select("sno", "sname").Where("sname =?", name).Get(student)
	//fmt.Println(student)
}
