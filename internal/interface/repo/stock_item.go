package repo

import (
	"context"
	"rent-product/internal/entity/stock_item"
)

type StockItemDB interface {
	GetStockItemAvailability(ctx context.Context, productID int64, stockItemList *[]*stock_item.StockItem) (err error)
}
