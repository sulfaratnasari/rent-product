package lostanimal

import (
	"context"
	"rent-product/internal/entity/order"
	"rent-product/internal/entity/product"
	"rent-product/internal/entity/stock_item"
	"rent-product/internal/interface/repo"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var mapMonthDays = map[int]int{
	1:  31,
	2:  29,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

type Usecase struct {
	ProductDB   repo.ProductDB
	OrderDB     repo.OrderDB
	StockItemDB repo.StockItemDB
}

func New(uc *Usecase) *Usecase {
	return uc
}

func (uc *Usecase) ProductAvailabilityList(ctx context.Context, productID int64, month, year int) (productAvailability product.ProductAvailability, err error) {
	stockItemList := []*stock_item.StockItem{}
	err = uc.StockItemDB.GetStockItemAvailability(ctx, productID, &stockItemList)
	if err != nil {
		return productAvailability, errors.Wrap(err, "usecase.ProductAvailabilityList.GetStockItemAvailability")
	}
	totalStockItem := len(stockItemList)

	orderList := []*order.Order{}
	err = uc.OrderDB.GetOrderListByProductID(ctx, productID, month, year, &orderList)
	if err != nil {
		return productAvailability, errors.Wrap(err, "usecase.ProductAvailabilityList.GetOrderListByProductID")
	}

	stockItemData := []stock_item.StockAvailability{}
	totalDays := mapMonthDays[month]

	// Create a WaitGroup to synchronize Goroutines
	var wg sync.WaitGroup

	// Create a channel to collect stock item data
	stockItemDataCh := make(chan stock_item.StockAvailability, totalDays)

	// Loop through each day in the month
	for i := 1; i <= totalDays; i++ {
		wg.Add(1) // Increment WaitGroup counter for each Goroutine

		go func(day int) {
			defer wg.Done() // Decrement WaitGroup counter when Goroutine completes

			dateToCheck := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			unavailableStock := checkUnavailableStock(orderList, dateToCheck)

			stockForCurrentDate := stock_item.StockAvailability{
				Date:           dateToCheck,
				AvailableStock: totalStockItem - unavailableStock,
			}

			// Send the stock item data to the channel
			stockItemDataCh <- stockForCurrentDate
		}(i)
	}

	// Close the channel when all Goroutines are done
	go func() {
		wg.Wait()
		close(stockItemDataCh)
	}()

	// Collect and process stock item data from the channel
	for stock := range stockItemDataCh {
		stockItemData = append(stockItemData, stock)
	}

	productAvailability.ProductID = productID
	productAvailability.Month = month
	productAvailability.Year = year
	productAvailability.TotalStockItem = totalStockItem
	productAvailability.StockItemData = stockItemData

	return productAvailability, nil
}

func checkUnavailableStock(orderList []*order.Order, date time.Time) (unavailableStock int) {

	for _, order := range orderList {
		startDate := order.StartDate
		endDate := order.EndDate

		if date.After(startDate) && date.Before(endDate) || date.Equal(startDate) || date.Equal(endDate) {
			unavailableStock++
		}
	}

	return unavailableStock
}

func (uc *Usecase) AddProduct(ctx context.Context, product product.Product) (err error) {

	_, err = uc.ProductDB.Add(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (uc *Usecase) ProductList(ctx context.Context) (productList []*product.Product, err error) {
	err = uc.ProductDB.GetProductList(ctx, &productList)
	if err != nil {
		return nil, err
	}
	return productList, nil
}
