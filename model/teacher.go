package model

type Teacher struct {
	Tno   string `gorm:"type:varchar(10);primary key;" xorm:"pk varchar(10)" json:"tno"` //教师编号
	Tname string `gorm:"type:varchar(20);" xorm:"varchar(20)" json:"tname"`              //教师姓名
}

func (table *Teacher) TableName() string {
	return "teacher"
}
