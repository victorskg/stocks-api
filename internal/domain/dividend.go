package stock

import (
	"github.com/google/uuid"
	"time"
)

type DType string

const (
	Income       DType = "income"
	Amortization DType = "amortization"
)

type Dividend struct {
	stockID     uuid.UUID
	value       float64
	baseDate    time.Time
	paymentDate time.Time
	dType       DType
}

// NewDividend TODO Add validations
func NewDividend(stockID uuid.UUID, value float64, baseDate time.Time, paymentDate time.Time, dType DType) *Dividend {
	return &Dividend{
		stockID:     stockID,
		value:       value,
		baseDate:    baseDate,
		paymentDate: paymentDate,
		dType:       dType,
	}
}

func (d Dividend) StockID() uuid.UUID {
	return d.stockID
}

func (d Dividend) Value() float64 {
	return d.value
}

func (d Dividend) BaseDate() time.Time {
	return d.baseDate
}

func (d Dividend) PaymentDate() time.Time {
	return d.paymentDate
}

func (d Dividend) DType() DType {
	return d.dType
}
