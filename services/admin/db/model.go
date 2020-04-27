package db

import (
	"github.com/jinzhu/gorm"
)

// type Model struct {
// 	ID        uint      `gorm:"primary_key; AUTO_INCREMENT"  json:"id,omitempty" example:"1"`
// 	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
// 	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
// DeletedAt *time.Time
// }

type TBL_USERS struct {
	gorm.Model
	Name  string `gorm:"unique_index; size:15"  json:"name,omitempty" example:"may"`
	Pwd   string `gorm:"default:'123456'; size:20"  json:"pwd,omitempty" example:"123"`
	Phone string `gorm:"size:11"  json:"phone,omitempty"`
	Email string `gorm:"size:25"  json:"email,omitempty"`
}

type TBL_VISIT_LOG struct {
	gorm.Model
	Host     string
	Method   string
	URL      string
	Header   string `gorm:"type:text"`
	Query    string `gorm:"type:text"`
	JSONForm string `gorm:"type:text"`
	JSONRaw  string `gorm:"type:text"`
}
