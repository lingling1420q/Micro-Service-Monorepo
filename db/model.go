package db

// User struct
type TBL_USERS struct {
	ID   uint   `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"size:255"` // string默认长度为255, 使用这种tag重设。
	Pwd  string `gorm:"default:'123'"`
}
