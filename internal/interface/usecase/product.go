package usecase

import (
	"context"
	"rent-product/internal/entity/product"
)

type ProductUC interface {
	ProductList(ctx context.Context) (productList []*product.Product, err error)
	AddProduct(ctx context.Context, product product.Product) (err error)
	ProductAvailabilityList(ctx context.Context, productID int64, month, year int) (productAvailability product.ProductAvailability, err error)
}
