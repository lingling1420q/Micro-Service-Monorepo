package defs

// User struct
type User struct {
	Name string `gorm:"size:255"` // string默认长度为255, 使用这种tag重设。
	Pwd  string `gorm:"default:'123'"`
	ID   int    `gorm:"AUTO_INCREMENT"`
}
