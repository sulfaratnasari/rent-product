package stock_item

import (
	"context"
	"rent-product/internal/entity/stock_item"
	"rent-product/lib/database/xorm"

	"github.com/pkg/errors"
)

type Conn struct {
	DB *xorm.Connection
}

func (conn *Conn) GetStockItemAvailability(ctx context.Context, productID int64, stockItemList *[]*stock_item.StockItem) (err error) {
	xormObj := conn.DB.Master.Context(ctx)

	xormObj.Where("product_id = ?", productID)
	err = xormObj.Find(stockItemList)
	if err != nil {
		return errors.Wrap(err, "repo.stock_item.GetStockItemAvailability")
	}
	return nil
}
