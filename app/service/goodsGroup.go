package service

import (
	"mugg/sumo/app/db"
	"mugg/sumo/model"
)

// GoodsGroup is
type GoodsGroup struct{}

// GetAll is
func (s GoodsGroup) GetAll() ([]model.GoodsGroup, error) {

	var (
		data []model.GoodsGroup
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
