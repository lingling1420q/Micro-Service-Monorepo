package db

import "time"

type TBL_USERS struct {
	ID    uint   `gorm:"primary_key; AUTO_INCREMENT"  json:"id,omitempty" example:"1"`
	Name  string `gorm:"unique_index; size:15"  json:"name,omitempty" example:"may"`
	Pwd   string `gorm:"default:'123456'; size:20"  json:"pwd,omitempty" example:"123"`
	Phone string `gorm:"size:11"  json:"phone,omitempty"`
	Email string `gorm:"size:25"  json:"email,omitempty"`
}

func (tbl_user *TBL_USERS) Insert() {
	gormConn := ConnGormMysql()
	gormConn.Create(&tbl_user)

}

type TBL_VISIT_LOG struct {
	Host      string
	Query     string `gorm:"type:text"`
	Body      string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (tvl *TBL_VISIT_LOG) Insert() {
	gormConn := ConnGormMysql()
	gormConn.Create(&tvl)

}
