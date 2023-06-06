package entities

import (
	domain "github.com/victorskg/stocks-api/internal/domain"
	"time"
)

import (
	"github.com/google/uuid"
)

type Dividend struct {
	ID          uuid.UUID    `bson:"_id,omitempty"`
	StockID     uuid.UUID    `bson:"stockID"`
	Value       float64      `bson:"value"`
	BaseDate    time.Time    `bson:"baseDate"`
	PaymentDate time.Time    `bson:"paymentDate"`
	DType       domain.DType `bson:"dType"`
}

func NewDividend(dividend domain.Dividend) *Dividend {
	return &Dividend{
		StockID:     dividend.StockID(),
		Value:       dividend.Value(),
		BaseDate:    dividend.BaseDate(),
		PaymentDate: dividend.PaymentDate(),
		DType:       dividend.DType(),
	}
}
