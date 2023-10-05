package order

import (
	"context"
	"rent-product/internal/entity/order"
	"rent-product/lib/database/xorm"

	"github.com/pkg/errors"
)

type Conn struct {
	DB *xorm.Connection
}

func (conn *Conn) GetOrderListByProductID(ctx context.Context, productID int64, month, year int, orderList *[]*order.Order) (err error) {
	xormObj := conn.DB.Master.Context(ctx)

	xormObj.Table("rp_order")
	xormObj.Select("rp_order.*")
	xormObj.Join("INNER", "rp_stock_item", "rp_order.stock_item_id=rp_stock_item.id")
	xormObj.Join("INNER", "rp_product", "rp_stock_item.product_id=rp_product.id")
	xormObj.Where("rp_product.id = ?", productID)

	err = xormObj.Find(orderList)
	if err != nil {
		return errors.Wrap(err, "repo.order.GetOrderListByProductID")
	}
	return nil
}
