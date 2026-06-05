package repository

import (
	"backend/internal/model"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

const maxPage = 10000

var ErrNotFound = errors.New("not found")

type ProductRepository interface {
	Create(ctx context.Context, p *model.Product) error
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	List(ctx context.Context, page, limit int) ([]model.Product, int, error)
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &productRepository{db: db}
}

func numericToDecimal(n pgtype.Numeric) decimal.Decimal {
	if !n.Valid {
		return decimal.Zero
	}
	intVal := n.Int
	if intVal == nil {
		intVal = new(big.Int)
	}
	return decimal.NewFromBigInt(intVal, n.Exp)
}

func decimalToNumeric(d decimal.Decimal) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   d.Coefficient(),
		Exp:   d.Exponent(),
		Valid: true,
	}
}

func (r *productRepository) Create(ctx context.Context, p *model.Product) error {
	if p.Name == "" {
		return fmt.Errorf("product name is required")
	}
	if p.Price.IsNegative() {
		return fmt.Errorf("product price cannot be negative")
	}
	if p.Stock < 0 {
		return fmt.Errorf("product stock cannot be negative")
	}

	query := `INSERT INTO products (category_id, name, sku, barcode, price, stock, unit, min_stock, is_active, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, true, NOW(), NOW()) RETURNING id`

	err := r.db.QueryRow(
		ctx,
		query,
		p.CategoryID, p.Name, p.SKU, p.Barcode, decimalToNumeric(p.Price), p.Stock, p.Unit, p.MinStock).Scan(&p.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("repository: duplicate sku or barcode: %w", err)
		}
		return fmt.Errorf("repository: failed to insert product: %w", err)
	}

	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	query := `SELECT id, category_id, name, sku, barcode, price, stock, unit, min_stock, is_active, created_at, updated_at
              FROM products WHERE id = $1`

	var p model.Product
	var priceNumeric pgtype.Numeric
	err := r.db.QueryRow(ctx, query, id).Scan(
		&p.ID, &p.CategoryID, &p.Name, &p.SKU, &p.Barcode, &priceNumeric,
		&p.Stock, &p.Unit, &p.MinStock, &p.IsActive, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: product not found", ErrNotFound)
		}
		return nil, fmt.Errorf("repository: failed to get product: %w", err)
	}

	p.Price = numericToDecimal(priceNumeric)
	return &p, nil
}

func (r *productRepository) List(ctx context.Context, page, limit int) ([]model.Product, int, error) {
	if page < 1 {
		page = 1
	}
	if page > maxPage {
		page = maxPage
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit
	var total int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM products").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("repository: failed to count products: %w", err)
	}
	query := `SELECT id, category_id, name, sku, barcode, price, stock, unit, min_stock, is_active, created_at, updated_at
	          FROM products ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("repository: failed to query products: %w", err)
	}
	defer rows.Close()
	var products []model.Product
	for rows.Next() {
		var p model.Product
		var priceNumeric pgtype.Numeric
		if err := rows.Scan(
			&p.ID, &p.CategoryID, &p.Name, &p.SKU, &p.Barcode,
			&priceNumeric, &p.Stock, &p.Unit, &p.MinStock, &p.IsActive,
			&p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("repository: failed to scan product: %w", err)
		}
		p.Price = numericToDecimal(priceNumeric)
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("repository: rows iteration error: %w", err)
	}
	return products, total, nil
}
