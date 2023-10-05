package stock_item

import "time"

type StockItem struct {
	ID        int64 `json:"id" xorm:"id pk" primary_key:"true"`
	ProductID int64 `json:"product_id" xorm:"product_id"`
}

type StockAvailability struct {
	Date           time.Time `json:"date"`
	AvailableStock int       `json:"available_stock"`
}
