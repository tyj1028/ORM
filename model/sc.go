package model

type SC struct {
	Sno   string  `gorm:"type:varchar(10);primary key" xorm:"varchar(10) pk" json:"sno"` //学号
	Cno   string  `gorm:"type:varchar(10);primary key" xorm:"varchar(10) pk" json:"cno"` //课程号
	Score float64 `gorm:"type:numeric(4,2)" xorm:"numeric(4,2)" json:"score"`            //学生成绩
}

func (table *SC) TableName() string {
	return "sc"
}
