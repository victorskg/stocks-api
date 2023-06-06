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

func NewStock(ticker string, name string, SType domain.SType, category string, subCategory string,
	administrator string, bookValue float32, patrimony float64, pvp float32) *Stock {
	return &Stock{
		Ticker:        ticker,
		Name:          name,
		SType:         SType,
		Category:      category,
		SubCategory:   subCategory,
		Administrator: administrator,
		BookValue:     bookValue,
		Patrimony:     patrimony,
		Pvp:           pvp,
	}
}
