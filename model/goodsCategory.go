package model

// GoodsCategory is 商品类型
type GoodsCategory struct {
	CategoryID        int64  `gorm:"column:category_id" json:"category_id" form:"category_id"`
	CategoryName      string `gorm:"column:category_name" json:"category_name" form:"category_name"`
	ShortName         string `gorm:"column:short_name" json:"short_name" form:"short_name"`
	Pid               int64  `gorm:"column:pid" json:"pid" form:"pid"`
	Level             int64  `gorm:"column:level" json:"level" form:"level"`
	IsVisible         int64  `gorm:"column:is_visible" json:"is_visible" form:"is_visible"`
	AttrID            int64  `gorm:"column:attr_id" json:"attr_id" form:"attr_id"`
	AttrName          string `gorm:"column:attr_name" json:"attr_name" form:"attr_name"`
	Keywords          string `gorm:"column:keywords" json:"keywords" form:"keywords"`
	Description       string `gorm:"column:description" json:"description" form:"description"`
	Sort              int64  `gorm:"column:sort" json:"sort" form:"sort"`
	CategoryPic       string `gorm:"column:category_pic" json:"category_pic" form:"category_pic"`
	PcCustomTemplate  string `gorm:"column:pc_custom_template" json:"pc_custom_template" form:"pc_custom_template"`
	WapCustomTemplate string `gorm:"column:wap_custom_template" json:"wap_custom_template" form:"wap_custom_template"`
}

//TableName is set
func (GoodsCategory) TableName() string {
	return "ns_goods_category"
}
