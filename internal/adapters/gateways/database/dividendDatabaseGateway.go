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
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dividendsCollection = "dividends"

type DividendDatabaseGateway struct {
	mapper       mappers.Mapper[entities.Dividend, domain.Dividend]
	stockGateway gateways.StockGateway
	collection   *mongo.Collection
}

func NewDividendDatabaseGateway(connection mongodb.Connection, stockGateway gateways.StockGateway) gateways.DividendGateway {
	return DividendDatabaseGateway{
		collection:   connection.Collection(dividendsCollection),
		stockGateway: stockGateway,
		mapper:       mappers.DividendMapper(),
	}
}

func (g DividendDatabaseGateway) SaveDividends(ticker string, dividends []domain.Dividend) ([]domain.Dividend, error) {
	stock, err := g.stockGateway.FindByTicker(ticker)
	if err != nil {
		return nil, err
	}

	documents := make([]interface{}, len(dividends))
	for i, dividend := range dividends {
		dividendEntity := g.mapper.FromDomainToEntity(dividend)
		dividendEntity.StockID = stock.Id()
		documents[i] = dividendEntity
	}

	_, err = g.collection.InsertMany(context.Background(), documents)
	if err != nil {
		fmt.Printf("Error inserting dividends of stock %s. Caused by: %s", ticker, err.Error())
		return nil, err
	}

	return dividends, nil
}

func (g DividendDatabaseGateway) FindByTicker(ticker string) ([]domain.Dividend, error) {
	stock, err := g.stockGateway.FindByTicker(ticker)
	if err != nil {
		return nil, err
	}

	opts := options.Find().SetSort(bson.D{{"baseDate", 1}})
	filter := bson.D{{"stockID", stock.Id()}}

	cursor, err := g.collection.Find(context.Background(), filter, opts)
	if err != nil {
		fmt.Printf("Error searching dividends of stock %s. Caused by: %s", ticker, err.Error())
		return nil, err
	}

	var results []entities.Dividend
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Printf("Error decoding dividends of stock %s. Caused by: %s", ticker, err.Error())
		return nil, err
	}

	dividendsDomain := make([]domain.Dividend, len(results))
	for i, result := range results {
		dividendsDomain[i] = *g.mapper.FromEntityToDomain(result)
	}

	return dividendsDomain, nil
}
