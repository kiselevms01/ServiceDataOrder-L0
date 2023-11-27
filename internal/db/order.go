package db

import (
	"fmt"

	"github.com/kiselevms01/wbProject_L0/model"
)

func (h *Handler) AutoMigrate() error {
	err := h.db.AutoMigrate(&model.Order{})
	if err != nil {
		return fmt.Errorf("error AutoMigrate: %v", err)
	}
	return nil
}

// Append to the Order table
func (h *Handler) AddDbOrder(order *model.Order) error {
	result := h.db.Create(&order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get all orders
func (r *Handler) GetAllDbOrder() ([]model.Order, error) {
	var allOrders []model.Order
	result := r.db.Find(&allOrders)

	if result.Error != nil {
		return nil, result.Error
	}

	return allOrders, nil
}
