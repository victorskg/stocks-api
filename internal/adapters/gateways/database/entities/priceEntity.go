package entities

import (
	"time"

	"github.com/google/uuid"
)

type Price struct {
	ID      uuid.UUID `bson:"_id,omitempty"`
	StockID uuid.UUID `bson:"stockId"`
	Date    time.Time `bson:"date"`
	Price   float64   `bson:"price"`
}
