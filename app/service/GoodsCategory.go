package service

import (
	"mugg/sumo/app/db"
	"mugg/sumo/model"
)

// GoodsCategory is
type GoodsCategory struct{}

// GetAll is
func (s GoodsCategory) GetAll() ([]model.GoodsCategory, error) {

	var (
		data []model.GoodsCategory
		err  error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Find(&data).Error; err != nil {
		tx.Rollback()
		return data, err
	}
	tx.Commit()

	return data, err

}
