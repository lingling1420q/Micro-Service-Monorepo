package db

import (
	"github.com/jinzhu/gorm"
)

type tblStruct interface {
	insert(*gorm.DB)
	//TODO: implement insert
	find(*gorm.DB) *gorm.DB
	delete(*gorm.DB)
}

func (u *TBL_USERS) insert(c *gorm.DB) {
	c.Create(&u)
}

func (u *TBL_USERS) find(c *gorm.DB) *gorm.DB {
	return c.Where(&u).Find(&TBL_USERS{})
}

func (u *TBL_USERS) delete(c *gorm.DB) {
	c.Where(&u).Delete(&TBL_USERS{})
}

func (vl *TBL_VISIT_LOG) insert(c *gorm.DB) {
	c.Create(&vl)
}

func (vl *TBL_VISIT_LOG) find(c *gorm.DB) *gorm.DB {
	return c.Where(&vl).Find(&TBL_VISIT_LOG{})
}

func (vl *TBL_VISIT_LOG) delete(c *gorm.DB) {
	c.Where(&vl).Delete(&TBL_VISIT_LOG{})
}

func InsertColume(t tblStruct) {
	gormConn := DB()
	t.insert(gormConn)
}

func FindColume(t tblStruct) *gorm.DB {
	gormConn := DB()
	return t.find(gormConn)
}

func Deletecolume(t tblStruct) {
	gormConn := DB()
	t.delete(gormConn)
}
