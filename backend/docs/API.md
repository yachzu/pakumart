# API Reference

Base URL: `http://localhost:8080`

## Authentication

Currently **not implemented** (planned for Phase 6). Endpoints marked "Protected" require a JWT `Authorization: Bearer <token>` header once auth is wired.

## Health

### `GET /health`

Check server and database connectivity.

**Response `200`**

```json
{ "status": "ok" }
```

**Response `500`**

```json
{ "status": "error", "error": "database connection failed" }
```

## Products (Planned)

### `GET /products`

List products with pagination.

| Query Param | Type | Default | Max | Description |
|-------------|------|---------|-----|-------------|
| `page` | int | 1 | 10000 | Page number |
| `limit` | int | 20 | 100 | Items per page |

**Response `200`**

```json
{
  "items": [
    {
      "id": 1,
      "name": "Beras 5kg",
      "category": "Sembako",
      "category_id": 2,
      "sku": "BR5K-001",
      "barcode": "8991234567890",
      "price": 12.50,
      "stock": 48,
      "unit": "kg",
      "min_stock": 10,
      "is_low_stock": false,
      "is_active": true
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 142,
    "total_page": 8
  }
}
```

### `GET /products/:id`

Get a single product.

**Response `200`** — Same shape as item in list.

**Response `404`**

```json
{ "error": "product not found" }
```

### `POST /products`

Create a product. Requires `admin` role (planned).

**Request body**

```json
{
  "name": "Beras 5kg",
  "category_id": 2,
  "sku": "BR5K-001",
  "barcode": "8991234567890",
  "price": 12.50,
  "stock": 50,
  "unit": "kg",
  "min_stock": 10
}
```

### `PUT /products/:id`

Update a product. Requires `admin` role (planned).

### `DELETE /products/:id`

Delete a product. Requires `admin` role (planned).

## Auth (Planned)

### `POST /auth/register`

Register a new user. Requires `admin` role.

### `POST /auth/login`

Authenticate and receive a JWT.

**Request body**

```json
{
  "username": "kasir1",
  "password": "secret123"
}
```

**Response `200`**

```json
{
  "token": "eyJhbGciOi...",
  "expired_at": 1735689600,
  "user": {
    "id": 1,
    "name": "Kasir Satu",
    "username": "kasir1",
    "role": "cashier",
    "is_active": true
  }
}
```

## Error Codes

| Code | Meaning |
|------|---------|
| `400` | Bad request — invalid input |
| `401` | Unauthorized — missing or invalid token |
| `403` | Forbidden — insufficient role |
| `404` | Not found |
| `405` | Method not allowed |
| `500` | Internal server error |

## Pagination

List endpoints return paginated responses with metadata:

```json
{
  "items": [],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 142,
    "total_page": 8
  }
}
```
