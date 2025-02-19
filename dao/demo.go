package dao

import (
	"tbwisk/public"
	"time"
)

//Area 1
type Area struct {
	ID       int       `json:"id" orm:"column(id);auto" description:"自增主键"`
	AreaName string    `json:"area_name" orm:"column(area_name);size(191)" description:"区域名称"`
	CityID   int       `json:"city_id" orm:"column(city_id)" description:"城市id"`
	UserID   int64     `json:"user_id" orm:"column(user_id)" description:"操作人"`
	UpdateAt time.Time `json:"update_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	CreateAt time.Time `json:"create_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	DeleteAt time.Time `json:"delete_at" orm:"column(delete_at);type(datetime)" description:"删除时间"`
}

//TableName 表明
func (f *Area) TableName() string {
	return "area"
}

//Find 根据id寻找
func (f *Area) Find(id string) ([]*Area, error) {
	var area []*Area
	err := public.GormPool.Where("id = ?", id).Find(&area).Error
	if err != nil {
		return nil, err
	}
	return area, nil
}

// func init() {
// 	public.GormPool.AutoMigrate(&Area{})
// }
