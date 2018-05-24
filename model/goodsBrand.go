package model

// GoodsBrand is
type GoodsBrand struct {
	BrandID           int64  `gorm:"column:brand_id" json:"brand_id" form:"brand_id"`
	ShopID            int64  `gorm:"column:shop_id" json:"shop_id" form:"shop_id"`
	BrandName         string `gorm:"column:brand_name" json:"brand_name" form:"brand_name"`
	BrandInitial      string `gorm:"column:brand_initial" json:"brand_initial" form:"brand_initial"`
	BrandPic          string `gorm:"column:brand_pic" json:"brand_pic" form:"brand_pic"`
	BrandRecommend    int64  `gorm:"column:brand_recommend" json:"brand_recommend" form:"brand_recommend"`
	Sort              int64  `gorm:"column:sort" json:"sort" form:"sort"`
	BrandCategoryName string `gorm:"column:brand_category_name" json:"brand_category_name" form:"brand_category_name"`
	CategoryIDArray   string `gorm:"column:category_id_array" json:"category_id_array" form:"category_id_array"`
	BrandAds          string `gorm:"column:brand_ads" json:"brand_ads" form:"brand_ads"`
	CategoryName      string `gorm:"column:category_name" json:"category_name" form:"category_name"`
	CategoryID1       int64  `gorm:"column:category_id_1" json:"category_id_1" form:"category_id_1"`
	CategoryID2       int64  `gorm:"column:category_id_2" json:"category_id_2" form:"category_id_2"`
	CategoryID3       int64  `gorm:"column:category_id_3" json:"category_id_3" form:"category_id_3"`
}

//TableName is set
func (GoodsBrand) TableName() string {
	return "ns_goods_brand"
}
