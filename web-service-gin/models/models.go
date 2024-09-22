package models

import (
	"time"
)

type Size string

const (
	SizeSmall  Size = "small"
	SizeMedium Size = "medium"
	SizeLarge  Size = "large"
)

type Status string

const (
	StatusCreated   Status = "created"
	StatusCancelled Status = "cancelled"
	StatusDispatched Status = "dispatched"
	StatusDelivered Status = "delivered"
	StatusProgress  Status = "progress"
)

type OrderItem struct {
	Product string `json:"product" binding:"required"`
	Size    Size    `json:"size" binding:"required,oneof=small medium large"`
	Quantity int    `json:"quantity" default:"1" binding:"gte=1,lte=1000000"`
}

type CreateOrder struct {
	OrderItems []OrderItem `json:"order_items" binding:"required,min=1"`
}

type GetOrder struct {
	ID        string     `json:"id"`
	Created   time.Time  `json:"created"`
	Status    Status     `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
}

type Error struct {
	Detail interface{} `json:"detail"`
}
