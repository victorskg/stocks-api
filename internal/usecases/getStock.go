package usecases

import (
	"github.com/victorskg/stocks-api/internal/domain"
	"github.com/victorskg/stocks-api/internal/usecases/gateways"
)

type GetStockUseCase struct {
	stockGateway    gateways.StockGateway
	priceGateway    gateways.PriceGateway
	dividendGateway gateways.DividendGateway
}

func NewGetStockUseCase(stockGateway gateways.StockGateway) GetStockUseCase {
	return GetStockUseCase{stockGateway: stockGateway}
}

func (u GetStockUseCase) Execute(ticker string) (*stock.Stock, error) {
	stockDomain, err := u.stockGateway.FindByTicker(ticker)
	if err != nil {
		return nil, err
	}

	stockPrices, err := u.priceGateway.FindByTicker(ticker)
	if err != nil {
		return nil, err
	}

	stockDividends, err := u.dividendGateway.FindByTicker(ticker)
	if err != nil {
		return nil, err
	}

	stockDomain.AddHistoricalPrices(stockPrices...)
	stockDomain.AddDividends(stockDividends...)

	return stockDomain, nil
}
