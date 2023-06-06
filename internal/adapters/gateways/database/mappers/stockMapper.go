package mappers

import (
	"github.com/victorskg/stocks-api/internal/adapters/gateways/database/entities"
	domain "github.com/victorskg/stocks-api/internal/domain"
)

type stockMapper struct {
}

func StockMapper() stockMapper {
	return stockMapper{}
}

func (m stockMapper) FromEntityToDomain(stockEntity entities.Stock) *domain.Stock {
	return domain.NewStock(stockEntity.ID, stockEntity.Ticker, stockEntity.Name, stockEntity.SType, stockEntity.Category,
		stockEntity.SubCategory, stockEntity.Administrator, stockEntity.BookValue, stockEntity.Patrimony, stockEntity.Pvp)
}

func (m stockMapper) FromDomainToEntity(stockDomain domain.Stock) *entities.Stock {
	return entities.NewStock(stockDomain)
}
