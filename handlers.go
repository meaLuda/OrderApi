package main

import (
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products := h.service.ListProducts()
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("productId")
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	product, err := h.service.GetProduct(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

type OrderHandler struct {
	orderService *OrderService
	promoService *PromoService
}

func NewOrderHandler(orderService *OrderService, promoService *PromoService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		promoService: promoService,
	}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Validate coupon if provided
	if req.CouponCode != "" {
		valid := h.promoService.ValidateCoupon(req.CouponCode)
		if !valid {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid coupon code"})
			return
		}
	}
	
	// Create order
	order, err := h.orderService.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, order)
}