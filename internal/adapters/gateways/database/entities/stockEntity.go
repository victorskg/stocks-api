package entities

import (
	"github.com/google/uuid"
	domain "github.com/victorskg/stocks-api/internal/domain"
)

type Stock struct {
	ID            uuid.UUID    `bson:"_id,omitempty"`
	Ticker        string       `bson:"ticker"`
	Name          string       `bson:"name"`
	SType         domain.SType `bson:"sType"`
	Category      string       `bson:"category"`
	SubCategory   string       `bson:"subCategory"`
	Administrator string       `bson:"administrator"`
	BookValue     float32      `bson:"bookValue"`
	Patrimony     float64      `bson:"patrimony"`
	Pvp           float32      `bson:"pvp"`
}

func NewStock(stock domain.Stock) *Stock {
	return &Stock{
		ID:            stock.Id(),
		Ticker:        stock.Ticker(),
		Name:          stock.Name(),
		SType:         stock.SType(),
		Category:      stock.Category(),
		SubCategory:   stock.SubCategory(),
		Administrator: stock.Administrator(),
		BookValue:     stock.BookValue(),
		Patrimony:     stock.Patrimony(),
		Pvp:           stock.Pvp(),
	}
}
