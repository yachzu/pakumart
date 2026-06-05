package model

import "time"

type StockLogType string

const (
	StockLogSale       StockLogType = "sale"
	StockLogRestock    StockLogType = "restock"
	StockLogAdjustment StockLogType = "adjustment"
	StockLogReturn     StockLogType = "return"
)

type StockLog struct {
	ID          uint         `json:"id"`
	ProductID   uint         `json:"product_id"`
	Product     *Product     `json:"product,omitempty"`
	Type        StockLogType `json:"type"`
	Quantity    int          `json:"quantity"`
	StockBefore int          `json:"stock_before"`
	StockAfter  int          `json:"stock_after"`
	ReferenceID *uint        `json:"reference_id,omitempty"`
	Note        string       `json:"note,omitempty"`
	CreatedBy   *uint        `json:"created_by,omitempty"`
	Creator     *User        `json:"creator,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
}
type RestockRequest struct {
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Note      string `json:"note"`
}
type AdjustStockRequest struct {
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Note      string `json:"note"`
}
type StockResponse struct {
	ProductID    uint   `json:"product_id"`
	ProductName  string `json:"product_name"`
	CategoryName string `json:"category_name"`
	CurrentStock int    `json:"current_stock"`
	Unit         string `json:"unit"`
	MinStock     int    `json:"min_stock"`
	IsLowStock   bool   `json:"is_low_stock"`
}
type StockLogResponse struct {
	ID          uint         `json:"id"`
	ProductID   uint         `json:"product_id"`
	ProductName string       `json:"product_name"`
	Type        StockLogType `json:"type"`
	Quantity    int          `json:"quantity"`
	StockBefore int          `json:"stock_before"`
	StockAfter  int          `json:"stock_after"`
	Note        string       `json:"note,omitempty"`
	CreatorName string       `json:"creator_name,omitempty"`
	CreatedAt   string       `json:"created_at"`
}

func (s *StockLog) ToResponse() StockLogResponse {
	var productName, creatorName string
	if s.Product != nil {
		productName = s.Product.Name
	}
	if s.Creator != nil {
		creatorName = s.Creator.Name
	}
	return StockLogResponse{
		ID:          s.ID,
		ProductID:   s.ProductID,
		ProductName: productName,
		Type:        s.Type,
		Quantity:    s.Quantity,
		StockBefore: s.StockBefore,
		StockAfter:  s.StockAfter,
		Note:        s.Note,
		CreatorName: creatorName,
		CreatedAt:   s.CreatedAt.Format(time.RFC3339),
	}
}
