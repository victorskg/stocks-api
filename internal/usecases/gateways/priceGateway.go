package gateways

import stock "github.com/victorskg/stocks-api/internal/domain"

type PriceGateway interface {
	FindByTicker(ticker string) ([]stock.Price, error)
}
