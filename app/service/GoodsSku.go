package service

import (
	"mugg/sumo/app/db"
	"mugg/sumo/model"
)

// GoodsSku is
type GoodsSku struct{}

// GetAll is
func (s GoodsSku) GetAll() ([]model.GoodsSku, error) {

	var (
		data []model.GoodsSku
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
