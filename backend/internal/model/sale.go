package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type PaymentMethod string

const (
	PaymentCash    PaymentMethod = "cash"
	PaymentQRIS    PaymentMethod = "qris"
	PaymentEWallet PaymentMethod = "e-wallet"
)

type Sale struct {
	ID            uint            `json:"id"`
	UserID        uint            `json:"user_id"`
	User          *User           `json:"user,omitempty"`
	TotalAmount   decimal.Decimal `json:"total_amount"`
	TotalItems    int             `json:"total_items"`
	PaymentMethod PaymentMethod   `json:"payment_method"`
	Notes         string          `json:"notes,omitempty"`
	CreatedAt     time.Time       `json:"created_at"`
	SaleItems     []SaleItem      `json:"items"`
}

type SaleItem struct {
	ID        uint            `json:"id"`
	SaleID    uint            `json:"sale_id"`
	ProductID uint            `json:"product_id"`
	Product   *Product        `json:"product,omitempty"`
	Quantity  decimal.Decimal `json:"quantity"`
	Unit      string          `json:"unit"`
	UnitPrice decimal.Decimal `json:"unit_price"`
	Subtotal  decimal.Decimal `json:"subtotal"`
	CreatedAt time.Time       `json:"created_at"`
}

type SaleItemRequest struct {
	ProductID uint            `json:"product_id"`
	Quantity  decimal.Decimal `json:"quantity"`
}

type CreateSaleRequest struct {
	Items         []SaleItemRequest `json:"items"`
	PaymentMethod PaymentMethod     `json:"payment_method"`
	Notes         string            `json:"notes"`
}

type SaleResponse struct {
	ID            uint             `json:"id"`
	UserID        uint             `json:"user_id"`
	CashierName   *string          `json:"cashier_name"`
	TotalAmount   decimal.Decimal  `json:"total_amount"`
	TotalItems    int              `json:"total_items"`
	PaymentMethod PaymentMethod    `json:"payment_method"`
	Notes         string           `json:"notes,omitempty"`
	CreatedAt     string           `json:"created_at"`
	Items         []SaleItemDetail `json:"items"`
}

type SaleItemDetail struct {
	ProductID   uint            `json:"product_id"`
	ProductName string          `json:"product_name"`
	Quantity    decimal.Decimal `json:"quantity"`
	Unit        string          `json:"unit"`
	UnitPrice   decimal.Decimal `json:"unit_price"`
	Subtotal    decimal.Decimal `json:"subtotal"`
}

func (s *Sale) ToResponse() SaleResponse {
	var cashierName *string
	if s.User != nil {
		name := s.User.Name
		cashierName = &name
	}

	items := make([]SaleItemDetail, len(s.SaleItems))
	for i, item := range s.SaleItems {
		productName := ""
		if item.Product != nil {
			productName = item.Product.Name
		}
		items[i] = SaleItemDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Unit:        item.Unit,
			UnitPrice:   item.UnitPrice,
			Subtotal:    item.Subtotal,
		}
	}

	return SaleResponse{
		ID:            s.ID,
		UserID:        s.UserID,
		CashierName:   cashierName,
		TotalAmount:   s.TotalAmount,
		TotalItems:    s.TotalItems,
		PaymentMethod: s.PaymentMethod,
		Notes:         s.Notes,
		CreatedAt:     s.CreatedAt.Format(time.RFC3339),
		Items:         items,
	}
}

type SaleListResponse struct {
	Items      []SaleResponse `json:"items"`
	Pagination PaginationMeta `json:"pagination"`
}
