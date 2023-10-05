package usecase

import (
	"context"
	"rent-product/internal/entity/product"
)

type ProductUC interface {
	ProductAvailabilityList(ctx context.Context, productID int64, month, year int) (productAvailability product.ProductAvailability, err error)
}
