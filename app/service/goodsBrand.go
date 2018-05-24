package service

import (
	"mugg/sumo/app/db"
	"mugg/sumo/model"
)

// GoodsBrand is
type GoodsBrand struct{}

// GetAll is
func (s GoodsBrand) GetAll() ([]model.GoodsBrand, error) {

	var (
		data []model.GoodsBrand
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
