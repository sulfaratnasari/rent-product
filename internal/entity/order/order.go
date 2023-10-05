package order

import "time"

type Order struct {
	ID          int64     `json:"id" xorm:"id pk" primary_key:"true"`
	StockItemID int64     `json:"stock_item_id" xorm:"stock_item_id"`
	StartDate   time.Time `json:"start_date" xorm:"start_date"`
	EndDate     time.Time `json:"end_date" xorm:"end_date"`
}
