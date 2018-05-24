package model

// GoodsSku is 商品标签
type GoodsSku struct {
	SkuID                int64   `gorm:"column:sku_id" json:"sku_id" form:"sku_id"`
	GoodsID              int64   `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	SkuName              string  `gorm:"column:sku_name" json:"sku_name" form:"sku_name"`
	AttrValueItems       string  `gorm:"column:attr_value_items" json:"attr_value_items" form:"attr_value_items"`
	AttrValueItemsFormat string  `gorm:"column:attr_value_items_format" json:"attr_value_items_format" form:"attr_value_items_format"`
	MarketPrice          float64 `gorm:"column:market_price" json:"market_price" form:"market_price"`
	Price                float64 `gorm:"column:price" json:"price" form:"price"`
	PromotePrice         float64 `gorm:"column:promote_price" json:"promote_price" form:"promote_price"`
	CostPrice            float64 `gorm:"column:cost_price" json:"cost_price" form:"cost_price"`
	Stock                int64   `gorm:"column:stock" json:"stock" form:"stock"`
	Picture              int64   `gorm:"column:picture" json:"picture" form:"picture"`
	Code                 string  `gorm:"column:code" json:"code" form:"code"`
	Qrcode               string  `gorm:"column:qrcode" json:"qrcode" form:"qrcode"`
	CreateDate           int64   `gorm:"column:create_date" json:"create_date" form:"create_date"`
	UpdateDate           int64   `gorm:"column:update_date" json:"update_date" form:"update_date"`
}

//TableName is set
func (GoodsSku) TableName() string {
	return "ns_goods_sku"
}
