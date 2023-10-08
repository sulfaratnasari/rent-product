package repo

import (
	"context"
	"rent-product/internal/entity/product"
)

type ProductDB interface {
	Add(ctx context.Context, product product.Product) (count int64, err error)
	GetProductList(ctx context.Context, productList *[]*product.Product) (err error)
}
