package service

import (
	"mugg/sumo/app/db"
	"mugg/sumo/model"
)

// Goods is
type Goods struct{}

// GetAll is
func (s Goods) GetAll() ([]model.Goods, error) {

	var (
		data []model.Goods
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
