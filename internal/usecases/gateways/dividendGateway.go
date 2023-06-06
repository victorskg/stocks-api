package gateways

import stock "github.com/victorskg/stocks-api/internal/domain"

type DividendGateway interface {
	SaveDividends(ticker string, dividends []stock.Dividend) ([]stock.Dividend, error)
	FindByTicker(ticker string) ([]stock.Dividend, error)
}
