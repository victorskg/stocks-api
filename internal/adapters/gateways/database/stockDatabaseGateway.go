package database

import (
	"context"
	"fmt"
	"github.com/victorskg/stocks-api/internal/adapters/gateways/database/entities"
	"github.com/victorskg/stocks-api/internal/adapters/gateways/database/mappers"
	domain "github.com/victorskg/stocks-api/internal/domain"
	"github.com/victorskg/stocks-api/internal/usecases/gateways"
	"github.com/victorskg/stocks-api/pkg/database/nosql/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "stocks"

type StockDatabaseGateway struct {
	mapper     mappers.Mapper[entities.Stock, domain.Stock]
	collection *mongo.Collection
}

func NewStockDatabaseGateway(connection mongodb.Connection) gateways.StockGateway {
	return StockDatabaseGateway{
		collection: connection.Collection(collection),
		mapper:     mappers.StockMapper(),
	}
}

func (g StockDatabaseGateway) SaveStocks(stocks []domain.Stock) ([]domain.Stock, error) {
	documents := make([]interface{}, len(stocks))
	for i, stock := range stocks {
		documents[i] = g.mapper.FromDomainToEntity(stock)
	}

	_, err := g.collection.InsertMany(context.Background(), documents)
	if err != nil {
		fmt.Printf("Error inserting stocks. Caused by: %s", err.Error())
		return nil, err
	}

	return stocks, nil
}

func (g StockDatabaseGateway) FindByTicker(ticker string) (*domain.Stock, error) {
	filter := bson.D{{"ticker", ticker}}

	var stockEntity entities.Stock
	err := g.collection.FindOne(context.Background(), filter).Decode(&stockEntity)
	if err != nil {
		return nil, err
	}

	return g.mapper.FromEntityToDomain(stockEntity), nil
}

func (g StockDatabaseGateway) UpdateStock(stock *domain.Stock) (*domain.Stock, error) {
	return nil, nil
}
