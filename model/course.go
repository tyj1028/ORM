package model

type Course struct {
	Cno   string `gorm:"type:varchar(10);primary_key;" xorm:"varchar(10) pk" json:"cno"` //课程号
	Cname string `gorm:"type:varchar(20);" xorm:"varchar(20)" json:"cname"`              //课程名
	Tno   string `gorm:"type:varchar(20);primary_key;" xorm:"varchar(20) pk" json:"tno"` //教师编号
}

func (table *Course) TableName() string {
	return "course"
}
