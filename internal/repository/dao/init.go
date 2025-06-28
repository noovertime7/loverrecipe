package dao

import "github.com/ego-component/egorm"

func InitTables(db *egorm.Component) error {
	// 自动迁移表结构
	err := db.AutoMigrate(
		&Dishes{},
		&DishType{},
		&User{},
	)
	return err
}
