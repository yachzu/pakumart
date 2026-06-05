CREATE TABLE public.users (
    id            BIGSERIAL PRIMARY KEY,
    name          VARCHAR NOT NULL,
    username      VARCHAR NOT NULL UNIQUE,
    password_hash VARCHAR NOT NULL,
    role          VARCHAR NOT NULL DEFAULT 'cashier',
    is_active     BOOLEAN DEFAULT true,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE public.categories (
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE public.products (
    id          BIGSERIAL PRIMARY KEY,
    category_id BIGINT REFERENCES public.categories(id),
    name        VARCHAR NOT NULL,
    sku         VARCHAR UNIQUE,
    barcode     VARCHAR UNIQUE,
    price       NUMERIC NOT NULL,
    stock       BIGINT DEFAULT 0,
    unit        VARCHAR NOT NULL,
    min_stock   BIGINT DEFAULT 5,
    is_active   BOOLEAN DEFAULT true,
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE public.sales (
    id             BIGSERIAL PRIMARY KEY,
    user_id        BIGINT NOT NULL REFERENCES public.users(id),
    total_amount   NUMERIC NOT NULL,
    total_items    BIGINT DEFAULT 0,
    payment_method VARCHAR DEFAULT 'cash',
    notes          TEXT,
    created_at     TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE public.sale_items (
    id         BIGSERIAL PRIMARY KEY,
    sale_id    BIGINT NOT NULL REFERENCES public.sales(id),
    product_id BIGINT NOT NULL REFERENCES public.products(id),
    quantity   NUMERIC NOT NULL,
    unit       VARCHAR,
    unit_price NUMERIC NOT NULL,
    subtotal   NUMERIC NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE public.stock_logs (
    id           BIGSERIAL PRIMARY KEY,
    product_id   BIGINT NOT NULL REFERENCES public.products(id),
    type         VARCHAR NOT NULL,
    quantity     BIGINT,
    stock_before BIGINT,
    stock_after  BIGINT,
    reference_id BIGINT,
    note         TEXT,
    created_by   BIGINT REFERENCES public.users(id),
    created_at   TIMESTAMPTZ DEFAULT NOW()
);
-- Indexes
CREATE INDEX idx_products_category ON public.products(category_id);
CREATE INDEX idx_products_sku ON public.products(sku);
CREATE INDEX idx_sales_user ON public.sales(user_id);
CREATE INDEX idx_sale_items_sale ON public.sale_items(sale_id);
CREATE INDEX idx_sale_items_product ON public.sale_items(product_id);
CREATE INDEX idx_stock_logs_product ON public.stock_logs(product_id);
CREATE INDEX idx_stock_logs_created ON public.stock_logs(created_at);
CREATE INDEX idx_products_name ON public.products(name);
CREATE INDEX idx_sales_created ON public.sales(created_at);
CREATE INDEX idx_users_username ON public.users(username);
