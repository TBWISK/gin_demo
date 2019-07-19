package public

import (
	"fmt"
	"tbwisk/common/lib"

	"github.com/jinzhu/gorm"
)

var (
	GormPool *gorm.DB
)

func InitMysql() error {
	dbpool, err := lib.GetGormPool("default")

	if err != nil {
		return err
	}
	GormPool = dbpool
	fmt.Println(GormPool.DB().Ping())

	return nil
}
