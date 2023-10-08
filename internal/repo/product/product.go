package product

import (
	"context"
	"rent-product/internal/entity/product"
	"rent-product/lib/database/xorm"

	"github.com/pkg/errors"
)

type Conn struct {
	DB *xorm.Connection
}

func (conn *Conn) Add(ctx context.Context, product product.Product) (count int64, err error) {
	xormObj := conn.DB.Master.Context(ctx)

	count, err = xormObj.Insert(product)

	if err != nil {
		return int64(0), errors.Wrap(err, "repo.product.Add")
	}
	return count, nil
}

func (conn *Conn) GetProductList(ctx context.Context, productList *[]*product.Product) (err error) {
	xormObj := conn.DB.Master.Context(ctx)

	err = xormObj.Find(productList)
	if err != nil {
		return errors.Wrap(err, "repo.product.GetProductList")
	}
	return nil
}
