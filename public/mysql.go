package public

import (
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
	return nil
}
