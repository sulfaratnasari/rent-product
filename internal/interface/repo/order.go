package repo

import (
	"context"
	"rent-product/internal/entity/order"
)

type OrderDB interface {
	GetOrderListByProductID(ctx context.Context, productID int64, month, year int, orderList *[]*order.Order) (err error)
}
