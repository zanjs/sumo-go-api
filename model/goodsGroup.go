package model

// GoodsGroup is 商品标签
type GoodsGroup struct {
	GroupID   int64  `gorm:"column:group_id" json:"group_id" form:"group_id"`
	ShopID    int64  `gorm:"column:shop_id" json:"shop_id" form:"shop_id"`
	GroupName string `gorm:"column:group_name" json:"group_name" form:"group_name"`
	Pid       int64  `gorm:"column:pid" json:"pid" form:"pid"`
	Level     int64  `gorm:"column:level" json:"level" form:"level"`
	IsVisible int64  `gorm:"column:is_visible" json:"is_visible" form:"is_visible"`
	GroupPic  string `gorm:"column:group_pic" json:"group_pic" form:"group_pic"`
	Sort      int64  `gorm:"column:sort" json:"sort" form:"sort"`
	GroupDec  string `gorm:"column:group_dec" json:"group_dec" form:"group_dec"`
}
