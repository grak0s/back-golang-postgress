package models

import "gorm.io/gorm"

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json: "last_name"`
	Name       string      `json:"name" gorm:"-"`
	Email      string      `json: "email"`
	Total      int         `json:"total" gorm:"-"`
	UpdateAt   string      `json:"updated_at"`
	CreateAt   string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items"  gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)
	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i, _ := range orders {

		var total float32 = 0

		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}
		orders[i].Name = orders[i].FirstName + " " + order.LastName
		orders[i].Total = int(total)
	}
	return orders
}
