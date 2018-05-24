package appCache

import (
	"fmt"
	"mugg/sumo/app/service"
	"mugg/sumo/model"

	"github.com/patrickmn/go-cache"
)

// Goods is
type Goods struct{}

// GetAll is
func (b Goods) GetAll() (interface{}, error) {

	var (
		data []model.Goods
		err  error
	)

	goods, has := Caches.Get("goods")

	if has {
		fmt.Println("has", goods)
		return goods, err
	}

	data, err = service.Goods{}.GetAll()
	fmt.Println("\b")
	fmt.Println(err)
	if err == nil {
		Caches.Set("goods", data, cache.DefaultExpiration)
	}

	return data, err
}
