package public

import (
	"fmt"

	"github.com/TBWISK/goconf"
	"github.com/jinzhu/gorm"
)

var (
	//GormPool gorm连接池
	GormPool *gorm.DB
)

//InitMysql mysql初始化
func InitMysql() error {
	GormPool = goconf.InitGorm("demo")
	// dbpool, err := lib.GetGormPool("demo")
	// fmt.Println("InitMysql", dbpool, err)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// GormPool = dbpool
	fmt.Println(GormPool.DB().Ping())

	return nil
}
