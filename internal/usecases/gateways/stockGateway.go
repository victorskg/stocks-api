package gateways

import stock "github.com/victorskg/stocks-api/internal/domain"

type StockGateway interface {
	SaveStocks(stocks []stock.Stock) ([]stock.Stock, error)
	FindByTicker(ticker string) (*stock.Stock, error)
	UpdateStock(stockDomain *stock.Stock) (*stock.Stock, error)
}
