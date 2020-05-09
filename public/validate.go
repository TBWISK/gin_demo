package public

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

var (
	//Uni 保存所有语言环境和翻译数据
	Uni *ut.UniversalTranslator
	//Validate 包含验证器设置和缓存
	Validate *validator.Validate
)

//InitValidate 校验初始化
func InitValidate() {
	en := en.New()
	zh := zh.New()
	zhtw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zhtw)
	Validate = validator.New()
}
