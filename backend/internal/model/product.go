package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID         uint            `json:"id"`
	CategoryID *uint           `json:"category_id"`
	Category   *Category       `json:"category"`
	Name       string          `json:"name"`
	SKU        string          `json:"sku,omitempty"`
	Barcode    string          `json:"barcode,omitempty"`
	Price      decimal.Decimal `json:"price"`
	Stock      int             `json:"stock"`
	Unit       string          `json:"unit"`
	MinStock   int             `json:"min_stock"`
	IsActive   bool            `json:"is_active"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

type ProductRequest struct {
	CategoryID *uint           `json:"category_id"`
	Name       string          `json:"name"`
	SKU        string          `json:"sku,omitempty"`
	Barcode    string          `json:"barcode,omitempty"`
	Price      decimal.Decimal `json:"price"`
	Stock      int             `json:"stock"`
	Unit       string          `json:"unit"`
	MinStock   int             `json:"min_stock"`
}

type ProductResponse struct {
	ID         uint            `json:"id"`
	Name       string          `json:"name"`
	Category   *string         `json:"category"`
	CategoryID *uint           `json:"category_id"`
	SKU        string          `json:"sku"`
	Barcode    string          `json:"barcode,omitempty"`
	Price      decimal.Decimal `json:"price,omitempty"`
	Stock      int             `json:"stock"`
	Unit       string          `json:"unit"`
	MinStock   int             `json:"min_stock"`
	IsLowStock bool            `json:"is_low_stock"`
	IsActive   bool            `json:"is_active"`
}

func (p *Product) ToResponse() ProductResponse {
	var categoryName *string
	if p.Category != nil {
		name := p.Category.Name
		categoryName = &name
	}
	return ProductResponse{
		ID:         p.ID,
		Name:       p.Name,
		Category:   categoryName,
		CategoryID: p.CategoryID,
		SKU:        p.SKU,
		Barcode:    p.Barcode,
		Price:      p.Price,
		Stock:      p.Stock,
		Unit:       p.Unit,
		MinStock:   p.MinStock,
		IsLowStock: p.Stock < p.MinStock,
		IsActive:   p.IsActive,
	}
}

type ProductListResponse struct {
	Items      []ProductResponse `json:"items"`
	Pagination PaginationMeta    `json:"pagination"`
}

type PaginationMeta struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}
