package main

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type ProductService struct {
	products map[int64]Product
}

func NewProductService() *ProductService {
	// some sample products
	products := map[int64]Product{
		1: {ID: "1", Name: "Chicken Waffle", Price: 12.99, Category: "Waffle"},
		2: {ID: "2", Name: "Classic Burger", Price: 9.99, Category: "Burger"},
	}

	return &ProductService{products: products}
}

func (s *ProductService) ListProducts() []Product {
	products := make([]Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}
	return products
}

func (s *ProductService) GetProduct(id int64) (*Product, error) {
	product, exists := s.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return &product, nil
}

type OrderService struct {
	productService *ProductService
}

func NewOrderService(productService *ProductService) *OrderService {
	return &OrderService{productService: productService}
}

func (s *OrderService) CreateOrder(req OrderRequest) (*Order, error) {
	// Validate all products exist and collect them
	var products []Product
	for _, item := range req.Items {
		productID, _ := strconv.ParseInt(item.ProductID, 10, 64)
		product, err := s.productService.GetProduct(productID)
		if err != nil {
			return nil, errors.New("invalid product ID: " + item.ProductID)
		}
		products = append(products, *product)
	}

	order := &Order{
		ID:       uuid.New().String(),
		Items:    req.Items,
		Products: products,
	}

	return order, nil
}

type PromoService struct {
	validCoupons map[string]bool
}

func NewPromoService() *PromoService {
	return &PromoService{
		validCoupons: loadValidCoupons(),
	}
}

func (s *PromoService) ValidateCoupon(code string) bool {
	// Length validation
	if len(code) < 8 || len(code) > 10 {
		return false
	}

	// Check if coupon exists in our valid coupons map
	return s.validCoupons[code]
}

func loadValidCoupons() map[string]bool {
	validCoupons := make(map[string]bool)

	// Load and process coupon files
	files := []string{"couponbase1.txt", "couponbase2.txt", "couponbase3.txt"}
	couponsInFiles := make(map[string]int)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		words := strings.Fields(string(content))
		for _, word := range words {
			if len(word) >= 8 && len(word) <= 10 {
				couponsInFiles[word]++
				if couponsInFiles[word] >= 2 {
					validCoupons[word] = true
				}
			}
		}
	}

	return validCoupons
}
