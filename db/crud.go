package db

import (
	"github.com/jinzhu/gorm"
)

type tblStruct interface {
	insert(*gorm.DB)
}

func (u *TBL_USERS) insert(c *gorm.DB) {
	c.Create(&u)
}

func (vl *TBL_VISIT_LOG) insert(c *gorm.DB) {
	c.Create(&vl)
}

func InsertColume(t tblStruct) {
	gormConn := ConnGormMysql()
	t.insert(gormConn)
}
