package main

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type OrderItem struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type OrderRequest struct {
	CouponCode string      `json:"couponCode"`
	Items      []OrderItem `json:"items" binding:"required,min=1"`
}

type Order struct {
	ID       string    `json:"id"`
	Items    []OrderItem `json:"items"`
	Products []Product  `json:"products"`
}