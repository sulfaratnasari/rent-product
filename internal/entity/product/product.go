package product

import "rent-product/internal/entity/stock_item"

type Product struct {
	ID          int64  `json:"id" xorm:"id pk" primary_key:"true"`
	Name        string `json:"name" xorm:"name"`
	Description string `json:"description" xorm:"description"`
}

type ProductAvailability struct {
	ProductID      int64                          `json:"product_id"`
	Month          int                            `json:"month"`
	Year           int                            `json:"year"`
	TotalStockItem int                            `json:"total_stock_item"`
	StockItemData  []stock_item.StockAvailability `json:"stock_item_data"`
}

type ProductParam struct {
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}
