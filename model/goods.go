package model

//Goods is 商品表
// type Goods struct {
// 	GoodsID   uint   `json:"goods_id" sql:"AUTO_INCREMENT" gorm:"unique_index;not null;unique;primary_key;column:goods_id"`
// 	GoodsName string `json:"goods_name"`
// }
type Goods struct {
	GoodsID             int64   `gorm:"column:goods_id" json:"goods_id" form:"goods_id"`
	GoodsName           string  `gorm:"column:goods_name" json:"goods_name" form:"goods_name"`
	ShopID              int64   `gorm:"column:shop_id" json:"shop_id" form:"shop_id"`
	CategoryID          int64   `gorm:"column:category_id" json:"category_id" form:"category_id"`
	CategoryID1         int64   `gorm:"column:category_id_1" json:"category_id_1" form:"category_id_1"`
	CategoryID2         int64   `gorm:"column:category_id_2" json:"category_id_2" form:"category_id_2"`
	CategoryID3         int64   `gorm:"column:category_id_3" json:"category_id_3" form:"category_id_3"`
	BrandID             int64   `gorm:"column:brand_id" json:"brand_id" form:"brand_id"`
	GroupIDArray        string  `gorm:"column:group_id_array" json:"group_id_array" form:"group_id_array"`
	PromotionType       int64   `gorm:"column:promotion_type" json:"promotion_type" form:"promotion_type"`
	PromoteID           int64   `gorm:"column:promote_id" json:"promote_id" form:"promote_id"`
	GoodsType           int64   `gorm:"column:goods_type" json:"goods_type" form:"goods_type"`
	MarketPrice         float64 `gorm:"column:market_price" json:"market_price" form:"market_price"`
	Price               float64 `gorm:"column:price" json:"price" form:"price"`
	PromotionPrice      float64 `gorm:"column:promotion_price" json:"promotion_price" form:"promotion_price"`
	CostPrice           float64 `gorm:"column:cost_price" json:"cost_price" form:"cost_price"`
	PointExchangeType   int64   `gorm:"column:point_exchange_type" json:"point_exchange_type" form:"point_exchange_type"`
	PointExchange       int64   `gorm:"column:point_exchange" json:"point_exchange" form:"point_exchange"`
	GivePoint           int64   `gorm:"column:give_point" json:"give_point" form:"give_point"`
	IsMemberDiscount    int64   `gorm:"column:is_member_discount" json:"is_member_discount" form:"is_member_discount"`
	ShippingFee         float64 `gorm:"column:shipping_fee" json:"shipping_fee" form:"shipping_fee"`
	ShippingFeeID       int64   `gorm:"column:shipping_fee_id" json:"shipping_fee_id" form:"shipping_fee_id"`
	Stock               int64   `gorm:"column:stock" json:"stock" form:"stock"`
	MaxBuy              int64   `gorm:"column:max_buy" json:"max_buy" form:"max_buy"`
	Clicks              int64   `gorm:"column:clicks" json:"clicks" form:"clicks"`
	MinStockAlarm       int64   `gorm:"column:min_stock_alarm" json:"min_stock_alarm" form:"min_stock_alarm"`
	Sales               int64   `gorm:"column:sales" json:"sales" form:"sales"`
	Collects            int64   `gorm:"column:collects" json:"collects" form:"collects"`
	Star                int64   `gorm:"column:star" json:"star" form:"star"`
	Evaluates           int64   `gorm:"column:evaluates" json:"evaluates" form:"evaluates"`
	Shares              int64   `gorm:"column:shares" json:"shares" form:"shares"`
	ProvinceID          int64   `gorm:"column:province_id" json:"province_id" form:"province_id"`
	CityID              int64   `gorm:"column:city_id" json:"city_id" form:"city_id"`
	Picture             int64   `gorm:"column:picture" json:"picture" form:"picture"`
	Keywords            string  `gorm:"column:keywords" json:"keywords" form:"keywords"`
	Introduction        string  `gorm:"column:introduction" json:"introduction" form:"introduction"`
	Description         string  `gorm:"column:description" json:"description" form:"description"`
	Qrcode              string  `gorm:"column:qrcode" json:"qrcode" form:"qrcode"`
	Code                string  `gorm:"column:code" json:"code" form:"code"`
	IsStockVisible      int64   `gorm:"column:is_stock_visible" json:"is_stock_visible" form:"is_stock_visible"`
	IsHot               int64   `gorm:"column:is_hot" json:"is_hot" form:"is_hot"`
	IsRecommend         int64   `gorm:"column:is_recommend" json:"is_recommend" form:"is_recommend"`
	IsNew               int64   `gorm:"column:is_new" json:"is_new" form:"is_new"`
	IsPreSale           int64   `gorm:"column:is_pre_sale" json:"is_pre_sale" form:"is_pre_sale"`
	IsBill              int64   `gorm:"column:is_bill" json:"is_bill" form:"is_bill"`
	State               int64   `gorm:"column:state" json:"state" form:"state"`
	Sort                int64   `gorm:"column:sort" json:"sort" form:"sort"`
	ImgIDArray          string  `gorm:"column:img_id_array" json:"img_id_array" form:"img_id_array"`
	SkuImgArray         string  `gorm:"column:sku_img_array" json:"sku_img_array" form:"sku_img_array"`
	MatchPoint          float64 `gorm:"column:match_point" json:"match_point" form:"match_point"`
	MatchRatio          float64 `gorm:"column:match_ratio" json:"match_ratio" form:"match_ratio"`
	RealSales           int64   `gorm:"column:real_sales" json:"real_sales" form:"real_sales"`
	GoodsAttributeID    int64   `gorm:"column:goods_attribute_id" json:"goods_attribute_id" form:"goods_attribute_id"`
	GoodsSpecFormat     string  `gorm:"column:goods_spec_format" json:"goods_spec_format" form:"goods_spec_format"`
	GoodsWeight         float64 `gorm:"column:goods_weight" json:"goods_weight" form:"goods_weight"`
	GoodsVolume         float64 `gorm:"column:goods_volume" json:"goods_volume" form:"goods_volume"`
	ShippingFeeType     int64   `gorm:"column:shipping_fee_type" json:"shipping_fee_type" form:"shipping_fee_type"`
	ExtendCategoryID    string  `gorm:"column:extend_category_id" json:"extend_category_id" form:"extend_category_id"`
	ExtendCategoryID1   string  `gorm:"column:extend_category_id_1" json:"extend_category_id_1" form:"extend_category_id_1"`
	ExtendCategoryID2   string  `gorm:"column:extend_category_id_2" json:"extend_category_id_2" form:"extend_category_id_2"`
	ExtendCategoryID3   string  `gorm:"column:extend_category_id_3" json:"extend_category_id_3" form:"extend_category_id_3"`
	SupplierID          int64   `gorm:"column:supplier_id" json:"supplier_id" form:"supplier_id"`
	SaleDate            int64   `gorm:"column:sale_date" json:"sale_date" form:"sale_date"`
	CreateTime          int64   `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime          int64   `gorm:"column:update_time" json:"update_time" form:"update_time"`
	MinBuy              int64   `gorm:"column:min_buy" json:"min_buy" form:"min_buy"`
	VirtualGoodsTypeID  int64   `gorm:"column:virtual_goods_type_id" json:"virtual_goods_type_id" form:"virtual_goods_type_id"`
	ProductionDate      int64   `gorm:"column:production_date" json:"production_date" form:"production_date"`
	ShelfLife           string  `gorm:"column:shelf_life" json:"shelf_life" form:"shelf_life"`
	GoodsVideoAddress   string  `gorm:"column:goods_video_address" json:"goods_video_address" form:"goods_video_address"`
	PcCustomTemplate    string  `gorm:"column:pc_custom_template" json:"pc_custom_template" form:"pc_custom_template"`
	WapCustomTemplate   string  `gorm:"column:wap_custom_template" json:"wap_custom_template" form:"wap_custom_template"`
	MaxUsePoint         int64   `gorm:"column:max_use_point" json:"max_use_point" form:"max_use_point"`
	IsOpenPresell       int64   `gorm:"column:is_open_presell" json:"is_open_presell" form:"is_open_presell"`
	PresellTime         int64   `gorm:"column:presell_time" json:"presell_time" form:"presell_time"`
	PresellDay          int64   `gorm:"column:presell_day" json:"presell_day" form:"presell_day"`
	PresellDeliveryType int64   `gorm:"column:presell_delivery_type" json:"presell_delivery_type" form:"presell_delivery_type"`
	PresellPrice        float64 `gorm:"column:presell_price" json:"presell_price" form:"presell_price"`
	GoodsUnit           string  `gorm:"column:goods_unit" json:"goods_unit" form:"goods_unit"`
}

//TableName is set User's table name to be `users`
func (Goods) TableName() string {
	return "ns_goods"
}
