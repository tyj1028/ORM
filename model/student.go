package model

type Student struct {
	Sno   string  `gorm:"type:varchar(10);primaryKey;" xorm:"pk varchar(10)" json:"sno"` //学号
	Sname string  `gorm:"type:varchar(20);" xorm:"varchar(20)" json:"sname"`             //姓名
	Sage  float64 `gorm:"type:numeric(2);" xorm:"numeric(2)" json:"sage"`                //年龄
	Ssex  string  `gorm:"type:varchar(5);" xorm:"varchar(5)" json:"ssex"`                //性别
}

func (table *Student) TableName() string {
	return "student"
}
