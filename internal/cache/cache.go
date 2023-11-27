package cache

import (
	"fmt"

	"github.com/kiselevms01/wbProject_L0/model"
)

type Cache struct {
	cache map[string]model.Order
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]model.Order),
	}
}

func (c *Cache) AddCacheOrder(order model.Order) {
	c.cache[order.OrderUid] = order
}

func (c *Cache) AddCacheOrders(orders []model.Order) error {
	for _, order := range orders {
		c.cache[order.OrderUid] = order
	}
	return nil
}

func (c *Cache) GetCacheOrder(orderUid string) (*model.Order, error) {
	order, boolVar := c.cache[orderUid]
	if boolVar {
		return &order, nil
	}
	return nil, fmt.Errorf("order not found")
}
