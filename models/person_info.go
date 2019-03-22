package models

import "time"

type PersonInfo struct {
	PersonId    int       `xorm:"not null pk autoincr INT(11)"`
	PlatformId  int       `xorm:"not null INT(11)"`
	PersonName  string    `xorm:"not null VARCHAR(60)"`
	Sex         string    `xorm:"VARCHAR(10)"`
	Age         int       `xorm:"INT(4)"`
	IdCard      string    `xorm:"not null unique VARCHAR(60)"`
	Address     string    `xorm:"VARCHAR(60)"`
	Party       string    `xorm:"VARCHAR(60)"`
	Phone       string    `xorm:"VARCHAR(50)"`
	Birthday    string    `xorm:"VARCHAR(50)"`
	CreatedTime time.Time `xorm:"not null DATETIME"`
	UpdatedTime time.Time `xorm:"not null DATETIME"`
	IsDeleted   int       `xorm:"default 0 INT(4)"`
}

func (m *PersonInfo) TableName() string {
	return "t_person_info"
}
