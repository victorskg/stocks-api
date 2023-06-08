package usecases

import (
	"github.com/victorskg/stocks-api/internal/domain"
	"github.com/victorskg/stocks-api/internal/usecases/gateways"
)

type SaveStocksUseCase struct {
	stockGateway gateways.StockGateway
}

func NewSaveStocksUseCase(stockGateway gateways.StockGateway) SaveStocksUseCase {
	return SaveStocksUseCase{stockGateway: stockGateway}
}

func (u SaveStocksUseCase) Execute(stocks ...stock.Stock) error {
	_, err := u.stockGateway.SaveStocks(stocks)
	return err
}
