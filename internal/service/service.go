package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kiselevms01/wbProject_L0/internal/cache"
	"github.com/kiselevms01/wbProject_L0/internal/db"
	"github.com/kiselevms01/wbProject_L0/model"
	"gorm.io/gorm"
)

type Client struct {
	database *gorm.DB
	cache    *cache.Cache
	newOrder *db.Handler
}

func NewOrderClient(database *gorm.DB) *Client {
	return &Client{
		database: database,
		cache:    cache.NewCache(),
		newOrder: db.NewDb(database),
	}
}

func (c *Client) Run() error {
	err := c.newOrder.AutoMigrate()
	if err != nil {
		return err
	}
	orders, err := c.newOrder.GetAllDbOrder()
	if err != nil {
		return err
	}
	c.cache.AddCacheOrders(orders)
	return nil
}

func (c *Client) AddOrder(data []byte) error {
	var order model.Order
	err := json.Unmarshal(data, &order)

	if err != nil {
		log.Printf("unmarshal error: %v", err)
	}

	fmt.Println("Get Order Uid: ", order.OrderUid)
	_, err = c.cache.GetCacheOrder(order.OrderUid)
	if err == nil {
		log.Printf("such an order already exists. Err: %v", err)
		return fmt.Errorf("such an order already exists")
	}

	err = c.newOrder.AddDbOrder(&order)
	if err != nil {
		log.Printf("error create order: %v", err)
		return err
	}
	c.cache.AddCacheOrder(order)
	return nil
}

func (c *Client) GetOrder(orderUid string) (*model.Order, error) {
	order, err := c.cache.GetCacheOrder(orderUid)
	if err != nil {
		// log.Printf("order not found %v", err)
		return nil, fmt.Errorf("order not found")
	}
	return order, nil
}
