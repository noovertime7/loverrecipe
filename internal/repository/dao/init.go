package dao

import (
	"github.com/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
)

func InitTables(db *egorm.Component) error {
	// 如果需要重新创建表，可以取消注释下面的代码
	// db.Migrator().DropTable(&User{}, &Dishes{}, &DishType{})

	// 自动迁移表结构
	err := db.AutoMigrate(
		&User{},
		&Dishes{},
		&DishType{},
	)

	if err != nil {
		elog.Error("数据库迁移失败", elog.FieldErr(err))
		return err
	}

	elog.Info("数据库表迁移成功")
	return nil
}
