package stock

import (
	"github.com/google/uuid"
	"time"
)

type SType string

const (
	Common     SType = "common"
	RealEstate SType = "real_estate"
)

type Price struct {
	date  time.Time
	price float64
}

type Stock struct {
	id               uuid.UUID
	ticker           string
	name             string
	sType            SType
	category         string
	subCategory      string
	administrator    string
	bookValue        float32
	patrimony        float64
	pvp              float32
	dividends        []Dividend
	historicalPrices []Price
}

func (s *Stock) Id() uuid.UUID {
	return s.id
}

func (s *Stock) Ticker() string {
	return s.ticker
}

func (s *Stock) Name() string {
	return s.name
}

func (s *Stock) SType() SType {
	return s.sType
}

func (s *Stock) Category() string {
	return s.category
}

func (s *Stock) SubCategory() string {
	return s.subCategory
}

func (s *Stock) Administrator() string {
	return s.administrator
}

func (s *Stock) BookValue() float32 {
	return s.bookValue
}

func (s *Stock) Patrimony() float64 {
	return s.patrimony
}

func (s *Stock) Pvp() float32 {
	return s.pvp
}

func (s *Stock) Dividends() []Dividend {
	return s.dividends
}

func (s *Stock) HistoricalPrices() []Price {
	return s.historicalPrices
}

// NewStockWithoutID TODO Add validations
func NewStockWithoutID(ticker string, name string, sType SType, category string, subCategory string,
	administrator string, bookValue float32, patrimony float64, pvp float32) *Stock {
	return &Stock{
		ticker:        ticker,
		name:          name,
		sType:         sType,
		category:      category,
		subCategory:   subCategory,
		administrator: administrator,
		bookValue:     bookValue,
		patrimony:     patrimony,
		pvp:           pvp,
	}
}

// NewStock TODO Add validations
func NewStock(id uuid.UUID, ticker string, name string, sType SType, category string, subCategory string,
	administrator string, bookValue float32, patrimony float64, pvp float32) *Stock {
	return &Stock{
		id:            id,
		ticker:        ticker,
		name:          name,
		sType:         sType,
		category:      category,
		subCategory:   subCategory,
		administrator: administrator,
		bookValue:     bookValue,
		patrimony:     patrimony,
		pvp:           pvp,
	}
}

func (s *Stock) AddHistoricalPrices(p ...Price) []Price {
	uniquePrices := make(map[Price]bool)
	for _, value := range s.historicalPrices {
		if !uniquePrices[value] {
			uniquePrices[value] = true
		}
	}

	var pricesToAdd []Price
	for _, value := range p {
		if !uniquePrices[value] {
			pricesToAdd = append(pricesToAdd, value)
			s.historicalPrices = append(s.historicalPrices, value)
		}
	}

	return pricesToAdd
}

func (s *Stock) AddDividends(d ...Dividend) []Dividend {
	uniqueDividends := make(map[Dividend]bool)
	for _, value := range s.dividends {
		if !uniqueDividends[value] {
			uniqueDividends[value] = true
		}
	}

	var dividendsToAdd []Dividend
	for _, value := range d {
		if !uniqueDividends[value] {
			dividendsToAdd = append(dividendsToAdd, value)
			s.dividends = append(s.dividends, value)
		}
	}

	return dividendsToAdd
}
