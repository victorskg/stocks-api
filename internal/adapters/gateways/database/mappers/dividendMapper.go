package mappers

import (
	"github.com/victorskg/stocks-api/internal/adapters/gateways/database/entities"
	domain "github.com/victorskg/stocks-api/internal/domain"
)

type dividendMapper struct {
}

func DividendMapper() dividendMapper {
	return dividendMapper{}
}

func (m dividendMapper) FromEntityToDomain(dividendEntity entities.Dividend) *domain.Dividend {
	return domain.NewDividend(dividendEntity.StockID, dividendEntity.Value, dividendEntity.BaseDate,
		dividendEntity.PaymentDate, dividendEntity.DType)
}

func (m dividendMapper) FromDomainToEntity(dividendDomain domain.Dividend) *entities.Dividend {
	return entities.NewDividend(dividendDomain)
}
