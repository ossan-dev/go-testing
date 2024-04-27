package user

import "gorm.io/gorm"

type User struct {
	ID   uint64
	Name string `gorm:"name_checker,name <> ''"`
}

func AddUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
